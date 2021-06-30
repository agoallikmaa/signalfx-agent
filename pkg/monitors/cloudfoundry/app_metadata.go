package cloudfoundry

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/sirupsen/logrus"
)

type ApplicationMetadataProvider struct {
	client *cfclient.Client
	logger logrus.FieldLogger
	cachedMetadata map[string]*ApplicationMetadata
}

type ApplicationMetadata struct {
	ApplicationName  string
	SpaceGuid        string
	SpaceName        string
	OrganizationGuid string
	OrganizationName string
}

func NewApplicationMetadataProvider(cloudControllerUrl string, uaaToken string, skipVerify bool, logger logrus.FieldLogger) *ApplicationMetadataProvider {
	if cloudControllerUrl == "" {
		logger.Info("Cloud Controller URL not provided, additional application metadata will not be provided")
		return nil
	}

	config := &cfclient.Config {
		ApiAddress: cloudControllerUrl,
		SkipSslValidation: skipVerify,
		// Skip "bearer " as this API wants the raw token
		Token:      uaaToken[7:],
	}

	client, clientError := cfclient.NewClient(config)
	if clientError != nil {
		logger.WithError(clientError).Error("Failed to connect to Cloud Controller API")
		return nil
	}

	logger.Infof("Using Cloud Controller at %s for fetching additional application metadata", cloudControllerUrl)

	return &ApplicationMetadataProvider{
		client: client,
		logger: logger,
		cachedMetadata: make(map[string]*ApplicationMetadata),
	}
}

func (p *ApplicationMetadataProvider) GetMetadata(applicationGuid string) *ApplicationMetadata {
	if entry, ok := p.cachedMetadata[applicationGuid]; ok {
		return entry
	}

	application, err := p.client.AppByGuid(applicationGuid)

	if err != nil {
		p.logger.WithError(err).Errorf("No application found for GUID %s", applicationGuid)
		p.cachedMetadata[applicationGuid] = nil
		return nil
	}

	entry := ApplicationMetadata{
		ApplicationName: application.Name,
		SpaceGuid: application.SpaceGuid,
		SpaceName: application.SpaceData.Entity.Name,
		OrganizationGuid: application.SpaceData.Entity.OrganizationGuid,
		OrganizationName: application.SpaceData.Entity.OrgData.Entity.Name,
	}

	p.logger.Debugf("Found metadata for application %s (%s)", applicationGuid, entry.ApplicationName)

	p.cachedMetadata[applicationGuid] = &entry
	return &entry
}
