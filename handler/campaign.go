package handler

import (
	"net/http"
	"strconv"

	"github.com/ekokurniawann/startup/campaign"
	"github.com/ekokurniawann/startup/helper"
	"github.com/gorilla/mux"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) FindCampaigns(w http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

	campaigns, err := h.service.FindCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Errro to find campaigns", http.StatusBadRequest, "error", nil)
		helper.RespondJSON(w, http.StatusBadRequest, response)
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "succes", campaign.FormatCampaigns(campaigns))
	helper.RespondJSON(w, http.StatusOK, response)
}

func (h *campaignHandler) FindCampaignByID(w http.ResponseWriter, r *http.Request) {
	var input campaign.GetCampaignDetailInput

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		response := helper.APIResponse("Invalid campaign ID", http.StatusBadRequest, "error", nil)
		helper.RespondJSON(w, http.StatusBadRequest, response)
		return
	}

	input.ID = id

	campaignDetail, err := h.service.FindCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		helper.RespondJSON(w, http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	helper.RespondJSON(w, http.StatusOK, response)

}
