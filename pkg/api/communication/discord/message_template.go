package discord

// discord only message template
type messageTemplate struct {
	Content string         `json:"content"`
	Embeds  []messageEmbed `json:"embeds"`
}

type messageEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
