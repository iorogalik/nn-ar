package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type OrganizationController struct {
	organizationService app.OrganizationService
}

func NewOrganizationController(os app.OrganizationService) OrganizationController {
	return OrganizationController{
		organizationService: os,
	}
}

func (c OrganizationController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org, err := requests.Bind(r, requests.OrganizationRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		org.UserId = user.Id
		org, err = c.organizationService.Save(org)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgDto resources.OrgDto
		Created(w, orgDto.DomainToDto(org))
	}
}

func (c OrganizationController) FindForUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		orgs, err := c.organizationService.FindForUser(user.Id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgsDto resources.OrgsDto
		response := orgsDto.DomainToDto(orgs)
		Success(w, response)
	}
}

func (c OrganizationController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)

		if org.UserId != user.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		var orgDto resources.OrgDto
		Success(w, orgDto.DomainToDto(org))
	}
}

func (c OrganizationController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org, err := requests.Bind(r, requests.OrganizationRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		organization := r.Context().Value(OrgKey).(domain.Organization)
		if organization.UserId != user.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		organization.Name = org.Name
		organization.Address = org.Address
		organization.City = org.City
		organization.Lat = org.Lat
		organization.Lon = org.Lon
		organization, err = c.organizationService.Update(organization)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgDto resources.OrgDto
		Success(w, orgDto.DomainToDto(organization))
	}
}
