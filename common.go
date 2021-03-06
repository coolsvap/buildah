package buildah

import (
	"io"

	cp "github.com/containers/image/copy"
	"github.com/containers/image/types"
)

func getCopyOptions(reportWriter io.Writer, sourceSystemContext *types.SystemContext, destinationSystemContext *types.SystemContext) *cp.Options {
	return &cp.Options{
		ReportWriter:   reportWriter,
		SourceCtx:      sourceSystemContext,
		DestinationCtx: destinationSystemContext,
	}
}

func getSystemContext(signaturePolicyPath string) *types.SystemContext {
	sc := &types.SystemContext{}
	if signaturePolicyPath != "" {
		sc.SignaturePolicyPath = signaturePolicyPath
	}
	return sc
}
