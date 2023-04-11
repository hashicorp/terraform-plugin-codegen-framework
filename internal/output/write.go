package output

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteDataSources(dataSources map[string][]byte, outputDir string) error {
	for k, v := range dataSources {
		filename := fmt.Sprintf("%s_data_source_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		//_, err = f.DataSourcesSchema(buf.Bytes())
		//if err != nil {
		//return err
		//}
	}

	return nil
}
