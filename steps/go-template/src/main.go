package main

import (
	"bytes"
	"context"
	"flag"
	"text/template"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/puppetlabs/relay-sdk-go/pkg/outputs"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
)

type TemplateSpec struct {
	Template   string                 `spec:"template"`
	Parameters map[string]interface{} `spec:"parameters"`
	Output     string                 `spec:"output"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	defaultMetadataSpecURL, err := taskutil.MetadataSpecURL()
	if err != nil {
		return err
	}

	// This seems like it could done better
	specURL := flag.String("spec-url", defaultMetadataSpecURL, "url to fetch spec from")
	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}
	spec := TemplateSpec{}

	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		return err
	}

	// Parameters must be an object when done this way - maybe we can detect for array types someway?
	params := spec.Parameters
	t, err := template.New("Render Template").Parse(spec.Template)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(spec.Parameters)+len(spec.Template)))
	if err := t.Execute(buf, params); err != nil {
		return errors.Wrap(err, "Could not fill out template")
	}

	oc, err := outputs.NewDefaultOutputsClientFromNebulaEnv()
	if err != nil {
		return err
	}

	ctx, cls := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cls()
	if err := oc.SetOutput(ctx, spec.Output, buf.String()); err != nil {
		return err
	}
	return nil
}
