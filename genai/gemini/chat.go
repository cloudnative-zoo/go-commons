package gemini

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

// GenerateCompletion sends a message to the generative model.
func (s *Service) GenerateCompletion(ctx context.Context, prompt string) (string, error) {
	model := s.client.GenerativeModel(s.model)
	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(response.Candidates) == 0 {
		if response.PromptFeedback.BlockReason == genai.BlockReasonSafety {
			for _, rating := range response.PromptFeedback.SafetyRatings {
				if rating.Blocked {
					return "", fmt.Errorf("completion blocked due to %v with probability %v", rating.Category.String(), rating.Probability.String())
				}
			}
		}
		return "", errors.New("no completion returned; unknown reason")
	}

	candidate := response.Candidates[0]
	var output string
	for _, part := range candidate.Content.Parts {
		switch content := part.(type) {
		case genai.Text:
			output += string(content)
			output += "\n"
		default:
			return "", fmt.Errorf("unexpected part type %T", content)
		}
	}

	if candidate.CitationMetadata != nil && len(candidate.CitationMetadata.CitationSources) > 0 {
		output += "Citations:\n"
		for _, source := range candidate.CitationMetadata.CitationSources {
			output += fmt.Sprintf("* %s, %s\n", *source.URI, source.License)
		}
	}

	return output, nil
}
