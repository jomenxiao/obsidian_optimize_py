package types

type TimeRequest struct {
	Timezone string `json:"timezone" description:"时区" required:"true"` // 使用 field tag 描述 inputschema
}
