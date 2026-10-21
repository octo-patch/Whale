package defaults

import "strings"

const (
	DefaultModel                 = "deepseek-v4-flash"
	ProModel                     = "deepseek-v4-pro"
	MiniMaxM31FlashModel         = "MiniMax-M3.1-Flash"
	MiniMaxM3Model               = "MiniMax-M3"
	DefaultReasoningEffort       = "high"
	DefaultThinkingEnabled       = true
	DefaultContextWindow         = 128_000
	DeepSeekV4ContextWindow      = 1_000_000
	MiniMaxM31FlashContextWindow = 1_000_000
	MiniMaxM3ContextWindow       = 1_000_000
	MiniMaxM27ContextWindow      = 204_800
	DefaultAutoCompactThreshold  = 0.85
	DefaultAgentCompactThreshold = 0.90
	// DefaultMaxToolIters is the ACP entrypoint's tool-iteration cap. It must
	// stay a finite backstop: the dynamic loop guards (storm rounds, progress
	// redundancy) cannot see mutating-argument-varying loops, so a cap is the
	// only guaranteed termination for that loop class. 300 sits far above any
	// recorded healthy turn (max 126) while bounding the invisible loops.
	DefaultMaxToolIters       = 300
	DefaultMemoryMaxChars     = 8000
	DefaultMemoryFileOrderCSV = "AGENTS.md,.claude/instructions.md,CLAUDE.md"
)

var supportedModels = []string{
	DefaultModel,
	ProModel,
	MiniMaxM31FlashModel,
	MiniMaxM3Model,
	"MiniMax-M2.7",
}

var defaultMemoryFileOrder = []string{
	"AGENTS.md",
	".claude/instructions.md",
	"CLAUDE.md",
}

func SupportedModels() []string {
	return append([]string(nil), supportedModels...)
}

func IsSupportedModel(model string) bool {
	m := strings.ToLower(strings.TrimSpace(model))
	for _, supported := range supportedModels {
		if strings.EqualFold(m, supported) {
			return true
		}
	}
	return false
}

func DefaultMemoryFileOrder() []string {
	return append([]string(nil), defaultMemoryFileOrder...)
}

func IsDeepSeekV4Model(model string) bool {
	m := strings.ToLower(strings.TrimSpace(model))
	return strings.Contains(m, DefaultModel) || strings.Contains(m, ProModel)
}

func IsMiniMaxM31FlashModel(model string) bool {
	return strings.EqualFold(strings.TrimSpace(model), MiniMaxM31FlashModel)
}

func IsMiniMaxM3Model(model string) bool {
	return strings.EqualFold(strings.TrimSpace(model), MiniMaxM3Model)
}

func IsMiniMaxM27Model(model string) bool {
	return strings.EqualFold(strings.TrimSpace(model), "MiniMax-M2.7")
}

// ContextWindowForModel returns the context window size in tokens for model.
func ContextWindowForModel(model string) int {
	if strings.TrimSpace(model) == "" {
		return DefaultContextWindow
	}
	if IsDeepSeekV4Model(model) {
		return DeepSeekV4ContextWindow
	}
	if IsMiniMaxM31FlashModel(model) {
		return MiniMaxM31FlashContextWindow
	}
	if IsMiniMaxM3Model(model) {
		return MiniMaxM3ContextWindow
	}
	if IsMiniMaxM27Model(model) {
		return MiniMaxM27ContextWindow
	}
	return DefaultContextWindow
}
