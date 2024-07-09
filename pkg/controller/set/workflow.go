package set

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghatm/pkg/edit"
)

func handleWorkflow(fs afero.Fs, file string, timeout int) error {
	b, err := afero.ReadFile(fs, file)
	if err != nil {
		return fmt.Errorf("read a file: %w", err)
	}
	after, err := edit.Edit(b, timeout)
	if err != nil {
		return fmt.Errorf("create a new workflow content: %w", err)
	}
	if after == nil {
		return nil
	}
	return writeWorkflow(fs, file, after)
}

func writeWorkflow(fs afero.Fs, file string, content []byte) error {
	stat, err := fs.Stat(file)
	if err != nil {
		return fmt.Errorf("get the workflow file stat: %w", err)
	}

	if err := afero.WriteFile(fs, file, content, stat.Mode()); err != nil {
		return fmt.Errorf("write the workflow file: %w", err)
	}
	return nil
}

func findWorkflows(fs afero.Fs) ([]string, error) {
	files, err := afero.Glob(fs, ".github/workflows/*.yml")
	if err != nil {
		return nil, fmt.Errorf("find .github/workflows/*.yml: %w", err)
	}
	files2, err := afero.Glob(fs, ".github/workflows/*.yaml")
	if err != nil {
		return nil, fmt.Errorf("find .github/workflows/*.yaml: %w", err)
	}
	return append(files, files2...), nil
}