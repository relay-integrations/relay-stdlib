package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/puppetlabs/relay-sdk-go/pkg/outputs"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
)

type TemplateSpec struct {
	Template string                 `spec:"template"`
	Data     map[string]interface{} `spec:"data"`
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

	t, err := template.New("template").Parse(spec.Template)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(spec.Data)+len(spec.Template)))
	if err := t.Execute(buf, spec.Data); err != nil {
		return fmt.Errorf("could not fill out template: %w", err)
	}

	oc, err := outputs.NewDefaultOutputsClientFromNebulaEnv()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := oc.SetOutput(ctx, "output", buf.String()); err != nil {
		return err
	}
	return nil
}
