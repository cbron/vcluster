package resources

import (
	synccontext "github.com/loft-sh/vcluster/pkg/controllers/syncer/context"
	"github.com/loft-sh/vcluster/pkg/mappings"
	"github.com/loft-sh/vcluster/pkg/mappings/generic"
	"github.com/loft-sh/vcluster/pkg/util/translate"
	corev1 "k8s.io/api/core/v1"
)

func CreateNamespacesMapper(ctx *synccontext.RegisterContext) (mappings.Mapper, error) {
	return generic.NewMapper(ctx, &corev1.Namespace{}, func(vName, _ string) string {
		return translate.Default.PhysicalNamespace(vName)
	})
}
