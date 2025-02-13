package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gopkg.in/yaml.v3"

	configv1alpha1 "github.com/stateful/runme/v3/internal/gen/proto/go/runme/config/v1alpha1"
)

// Config is a flatten configuration of runme.yaml. The purpose of it is to
// unify all the different configuration versions into a single struct.
type Config struct {
	// Dir- or git-based project fields.
	DisableGitignore bool
	IgnorePaths      []string
	FindRepoUpward   bool
	ProjectDir       string

	// Filemode fields.
	Filename string

	// Environment variable fields.
	EnvSourceFiles []string
	UseSystemEnv   bool

	Filters []*Filter

	// Log related fields.
	LogEnabled bool
	LogPath    string
	LogVerbose bool

	// Server related fields.
	ServerAddress     string
	ServerTLSEnabled  bool
	ServerTLSCertFile string
	ServerTLSKeyFile  string

	// Kernel related fields.
	Kernels []Kernel
}

func ParseYAML(data []byte) (*Config, error) {
	version, err := parseVersionFromYAML(data)
	if err != nil {
		return nil, err
	}
	switch version {
	case "v1alpha1":
		cfg, err := parseYAMLv1alpha1(data)
		if err != nil {
			return nil, err
		}

		if err := validateProto(cfg); err != nil {
			return nil, errors.Wrap(err, "failed to validate v1alpha1 config")
		}

		config, err := configV1alpha1ToConfig(cfg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert v1alpha1 config")
		}

		if err := validateConfig(config); err != nil {
			return nil, errors.Wrap(err, "failed to validate config")
		}

		return config, nil
	default:
		return nil, errors.Errorf("unknown version: %s", version)
	}
}

type versionOnly struct {
	Version string `yaml:"version"`
}

func parseVersionFromYAML(data []byte) (string, error) {
	var result versionOnly

	if err := yaml.Unmarshal(data, &result); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal version")
	}

	return result.Version, nil
}

func parseYAMLv1alpha1(data []byte) (*configv1alpha1.Config, error) {
	mmap := make(map[string]any)

	if err := yaml.Unmarshal(data, &mmap); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml")
	}

	delete(mmap, "version")

	// In order to properly handle JSON-related field options like `json_name`,
	// the YAML data is first marshaled to JSON and then unmarshaled to a proto message
	// using the protojson package.
	configJSONRaw, err := json.Marshal(mmap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal yaml to json")
	}

	var cfg configv1alpha1.Config
	if err := protojson.Unmarshal(configJSONRaw, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json to proto")
	}
	return &cfg, nil
}

func configV1alpha1ToConfig(c *configv1alpha1.Config) (*Config, error) {
	project := c.GetProject()
	log := c.GetLog()

	var filters []*Filter
	for _, f := range c.GetFilters() {
		filters = append(filters, &Filter{
			Type:      f.GetType().String(),
			Condition: f.GetCondition(),
		})
	}

	var kernels []Kernel
	for _, k := range c.GetKernels() {
		msg, err := k.UnmarshalNew()
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal kernel")
		}

		switch k := msg.(type) {
		case *configv1alpha1.Config_LocalKernel:
			kernels = append(kernels, &LocalKernel{
				Name: k.GetName(),
			})
		case *configv1alpha1.Config_DockerKernel:
			kernels = append(kernels, &DockerKernel{
				Build: struct {
					Context    string
					Dockerfile string
				}{
					Context:    k.GetBuild().GetContext(),
					Dockerfile: k.GetBuild().GetDockerfile(),
				},
				Image: k.GetImage(),
				Name:  k.GetName(),
			})
		default:
			return nil, errors.Errorf("unknown kernel type: %s", k.ProtoReflect().Type().Descriptor().FullName())
		}
	}

	cfg := &Config{
		ProjectDir:       project.GetDir(),
		FindRepoUpward:   project.GetFindRepoUpward(),
		IgnorePaths:      project.GetIgnorePaths(),
		DisableGitignore: project.GetDisableGitignore(),

		Filename: c.GetFilename(),

		UseSystemEnv:   c.GetEnv().GetUseSystemEnv(),
		EnvSourceFiles: c.GetEnv().GetSources(),

		Filters: filters,

		LogEnabled: log.GetEnabled(),
		LogPath:    log.GetPath(),
		LogVerbose: log.GetVerbose(),

		ServerAddress:     c.GetServer().GetAddress(),
		ServerTLSEnabled:  c.GetServer().GetTls().GetEnabled(),
		ServerTLSCertFile: c.GetServer().GetTls().GetCertFile(),
		ServerTLSKeyFile:  c.GetServer().GetTls().GetKeyFile(),

		Kernels: kernels,
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	if err := validateProjectDir(cfg, cwd); err != nil {
		return errors.Wrap(err, "failed to validate project dir")
	}

	if err := validateFilename(cfg, cwd); err != nil {
		return errors.Wrap(err, "failed to validate filename")
	}

	return nil
}

func validateProjectDir(cfg *Config, cwd string) error {
	rel, err := filepath.Rel(cwd, filepath.Join(cwd, cfg.ProjectDir))
	if err != nil {
		return errors.WithStack(err)
	}
	if strings.HasPrefix(rel, "..") {
		return errors.New("outside of current working directory")
	}

	return nil
}

func validateFilename(cfg *Config, cwd string) error {
	rel, err := filepath.Rel(cwd, filepath.Join(cwd, cfg.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	if strings.HasPrefix(rel, "..") {
		return errors.New("outside of current working directory")
	}

	return nil
}

func validateProto(m protoreflect.ProtoMessage) error {
	v, err := protovalidate.New()
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(v.Validate(m))
}
