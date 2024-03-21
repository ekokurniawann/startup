package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	imageUrl := ""
	if (len(campaign.CampaignImages)) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	formattedCampaign := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         imageUrl,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	return formattedCampaign
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignResponseData struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageUrl         string   `json:"image_url"`
	GoalAmount       int      `json:"goal_amount"`
	CurrentAmount    int      `json:"current_amount"`
	UserID           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
	User             struct {
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
	} `json:"user"`
	Images []struct {
		ImageUrl  string `json:"image_url"`
		IsPrimary bool   `json:"is_primary"`
	} `json:"images"`
}

func FormatCampaignDetail(campaign Campaign) CampaignResponseData {
	imageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	perks := strings.Split(campaign.Perks, ",")
	for i, perk := range perks {
		perks[i] = strings.TrimSpace(perk)
	}

	var images []struct {
		ImageUrl  string `json:"image_url"`
		IsPrimary bool   `json:"is_primary"`
	}
	for _, img := range campaign.CampaignImages {
		isPrimary := img.IsPrimary == 1
		images = append(images, struct {
			ImageUrl  string `json:"image_url"`
			IsPrimary bool   `json:"is_primary"`
		}{
			ImageUrl:  img.FileName,
			IsPrimary: isPrimary,
		})
	}

	responseData := CampaignResponseData{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageUrl:         imageUrl,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Perks:            perks,
		User: struct {
			Name     string `json:"name"`
			ImageUrl string `json:"image_url"`
		}{
			Name:     campaign.User.Name,
			ImageUrl: campaign.User.AvatarFileName,
		},
		Images: images,
	}

	return responseData
}
