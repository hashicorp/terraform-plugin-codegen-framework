package provider_apigw

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"os/exec"
	"time"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/common"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/conn"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ApigwProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{}
}

type ApigwModel struct {
}

