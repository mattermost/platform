// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package timerlayer

import (
	"context"
	timemodule "time"

	"github.com/mattermost/mattermost-server/v5/einterfaces"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
)

type TimerLayer struct {
	store.Store
	Metrics                   einterfaces.MetricsInterface
	AuditStore                store.AuditStore
	BotStore                  store.BotStore
	ChannelStore              store.ChannelStore
	ChannelMemberHistoryStore store.ChannelMemberHistoryStore
	ClusterDiscoveryStore     store.ClusterDiscoveryStore
	CommandStore              store.CommandStore
	CommandWebhookStore       store.CommandWebhookStore
	ComplianceStore           store.ComplianceStore
	EmojiStore                store.EmojiStore
	FileInfoStore             store.FileInfoStore
	GroupStore                store.GroupStore
	JobStore                  store.JobStore
	LicenseStore              store.LicenseStore
	LinkMetadataStore         store.LinkMetadataStore
	OAuthStore                store.OAuthStore
	PluginStore               store.PluginStore
	PostStore                 store.PostStore
	PreferenceStore           store.PreferenceStore
	ReactionStore             store.ReactionStore
	RoleStore                 store.RoleStore
	SchemeStore               store.SchemeStore
	SessionStore              store.SessionStore
	StatusStore               store.StatusStore
	SystemStore               store.SystemStore
	TeamStore                 store.TeamStore
	TermsOfServiceStore       store.TermsOfServiceStore
	TokenStore                store.TokenStore
	UserStore                 store.UserStore
	UserAccessTokenStore      store.UserAccessTokenStore
	UserTermsOfServiceStore   store.UserTermsOfServiceStore
	WebhookStore              store.WebhookStore
}

func (s *TimerLayer) Audit() store.AuditStore {
	return s.AuditStore
}

func (s *TimerLayer) Bot() store.BotStore {
	return s.BotStore
}

func (s *TimerLayer) Channel() store.ChannelStore {
	return s.ChannelStore
}

func (s *TimerLayer) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.ChannelMemberHistoryStore
}

func (s *TimerLayer) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.ClusterDiscoveryStore
}

func (s *TimerLayer) Command() store.CommandStore {
	return s.CommandStore
}

func (s *TimerLayer) CommandWebhook() store.CommandWebhookStore {
	return s.CommandWebhookStore
}

func (s *TimerLayer) Compliance() store.ComplianceStore {
	return s.ComplianceStore
}

func (s *TimerLayer) Emoji() store.EmojiStore {
	return s.EmojiStore
}

func (s *TimerLayer) FileInfo() store.FileInfoStore {
	return s.FileInfoStore
}

func (s *TimerLayer) Group() store.GroupStore {
	return s.GroupStore
}

func (s *TimerLayer) Job() store.JobStore {
	return s.JobStore
}

func (s *TimerLayer) License() store.LicenseStore {
	return s.LicenseStore
}

func (s *TimerLayer) LinkMetadata() store.LinkMetadataStore {
	return s.LinkMetadataStore
}

func (s *TimerLayer) OAuth() store.OAuthStore {
	return s.OAuthStore
}

func (s *TimerLayer) Plugin() store.PluginStore {
	return s.PluginStore
}

func (s *TimerLayer) Post() store.PostStore {
	return s.PostStore
}

func (s *TimerLayer) Preference() store.PreferenceStore {
	return s.PreferenceStore
}

func (s *TimerLayer) Reaction() store.ReactionStore {
	return s.ReactionStore
}

func (s *TimerLayer) Role() store.RoleStore {
	return s.RoleStore
}

func (s *TimerLayer) Scheme() store.SchemeStore {
	return s.SchemeStore
}

func (s *TimerLayer) Session() store.SessionStore {
	return s.SessionStore
}

func (s *TimerLayer) Status() store.StatusStore {
	return s.StatusStore
}

func (s *TimerLayer) System() store.SystemStore {
	return s.SystemStore
}

func (s *TimerLayer) Team() store.TeamStore {
	return s.TeamStore
}

func (s *TimerLayer) TermsOfService() store.TermsOfServiceStore {
	return s.TermsOfServiceStore
}

func (s *TimerLayer) Token() store.TokenStore {
	return s.TokenStore
}

func (s *TimerLayer) User() store.UserStore {
	return s.UserStore
}

func (s *TimerLayer) UserAccessToken() store.UserAccessTokenStore {
	return s.UserAccessTokenStore
}

func (s *TimerLayer) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.UserTermsOfServiceStore
}

func (s *TimerLayer) Webhook() store.WebhookStore {
	return s.WebhookStore
}

type TimerLayerAuditStore struct {
	store.AuditStore
	Root *TimerLayer
}

type TimerLayerBotStore struct {
	store.BotStore
	Root *TimerLayer
}

type TimerLayerChannelStore struct {
	store.ChannelStore
	Root *TimerLayer
}

type TimerLayerChannelMemberHistoryStore struct {
	store.ChannelMemberHistoryStore
	Root *TimerLayer
}

type TimerLayerClusterDiscoveryStore struct {
	store.ClusterDiscoveryStore
	Root *TimerLayer
}

type TimerLayerCommandStore struct {
	store.CommandStore
	Root *TimerLayer
}

type TimerLayerCommandWebhookStore struct {
	store.CommandWebhookStore
	Root *TimerLayer
}

type TimerLayerComplianceStore struct {
	store.ComplianceStore
	Root *TimerLayer
}

type TimerLayerEmojiStore struct {
	store.EmojiStore
	Root *TimerLayer
}

type TimerLayerFileInfoStore struct {
	store.FileInfoStore
	Root *TimerLayer
}

type TimerLayerGroupStore struct {
	store.GroupStore
	Root *TimerLayer
}

type TimerLayerJobStore struct {
	store.JobStore
	Root *TimerLayer
}

type TimerLayerLicenseStore struct {
	store.LicenseStore
	Root *TimerLayer
}

type TimerLayerLinkMetadataStore struct {
	store.LinkMetadataStore
	Root *TimerLayer
}

type TimerLayerOAuthStore struct {
	store.OAuthStore
	Root *TimerLayer
}

type TimerLayerPluginStore struct {
	store.PluginStore
	Root *TimerLayer
}

type TimerLayerPostStore struct {
	store.PostStore
	Root *TimerLayer
}

type TimerLayerPreferenceStore struct {
	store.PreferenceStore
	Root *TimerLayer
}

type TimerLayerReactionStore struct {
	store.ReactionStore
	Root *TimerLayer
}

type TimerLayerRoleStore struct {
	store.RoleStore
	Root *TimerLayer
}

type TimerLayerSchemeStore struct {
	store.SchemeStore
	Root *TimerLayer
}

type TimerLayerSessionStore struct {
	store.SessionStore
	Root *TimerLayer
}

type TimerLayerStatusStore struct {
	store.StatusStore
	Root *TimerLayer
}

type TimerLayerSystemStore struct {
	store.SystemStore
	Root *TimerLayer
}

type TimerLayerTeamStore struct {
	store.TeamStore
	Root *TimerLayer
}

type TimerLayerTermsOfServiceStore struct {
	store.TermsOfServiceStore
	Root *TimerLayer
}

type TimerLayerTokenStore struct {
	store.TokenStore
	Root *TimerLayer
}

type TimerLayerUserStore struct {
	store.UserStore
	Root *TimerLayer
}

type TimerLayerUserAccessTokenStore struct {
	store.UserAccessTokenStore
	Root *TimerLayer
}

type TimerLayerUserTermsOfServiceStore struct {
	store.UserTermsOfServiceStore
	Root *TimerLayer
}

type TimerLayerWebhookStore struct {
	store.WebhookStore
	Root *TimerLayer
}

func (s *TimerLayerAuditStore) Get(user_id string, offset int, limit int) (model.Audits, error) {
	start := timemodule.Now()

	result, err := s.AuditStore.Get(user_id, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("AuditStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerAuditStore) PermanentDeleteByUser(userId string) error {
	start := timemodule.Now()

	err := s.AuditStore.PermanentDeleteByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("AuditStore.PermanentDeleteByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerAuditStore) Save(audit *model.Audit) error {
	start := timemodule.Now()

	err := s.AuditStore.Save(audit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("AuditStore.Save", success, elapsed)
	}
	return err
}

func (s *TimerLayerBotStore) Get(userId string, includeDeleted bool) (*model.Bot, error) {
	start := timemodule.Now()

	result, err := s.BotStore.Get(userId, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("BotStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerBotStore) GetAll(options *model.BotGetOptions) ([]*model.Bot, error) {
	start := timemodule.Now()

	result, err := s.BotStore.GetAll(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("BotStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerBotStore) PermanentDelete(userId string) error {
	start := timemodule.Now()

	err := s.BotStore.PermanentDelete(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("BotStore.PermanentDelete", success, elapsed)
	}
	return err
}

func (s *TimerLayerBotStore) Save(bot *model.Bot) (*model.Bot, error) {
	start := timemodule.Now()

	result, err := s.BotStore.Save(bot)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("BotStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerBotStore) Update(bot *model.Bot) (*model.Bot, error) {
	start := timemodule.Now()

	result, err := s.BotStore.Update(bot)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("BotStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.AnalyticsDeletedTypeCount(teamId, channelType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.AnalyticsDeletedTypeCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) AnalyticsTypeCount(teamId string, channelType string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.AnalyticsTypeCount(teamId, channelType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.AnalyticsTypeCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) AutocompleteInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.AutocompleteInTeam(teamId, term, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.AutocompleteInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.AutocompleteInTeamForSearch(teamId, userId, term, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.AutocompleteInTeamForSearch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) ClearAllCustomRoleAssignments() *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.ClearAllCustomRoleAssignments()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.ClearAllCustomRoleAssignments", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) ClearCaches() {
	start := timemodule.Now()

	s.ChannelStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) ClearSidebarOnTeamLeave(userId string, teamId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.ClearSidebarOnTeamLeave(userId, teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.ClearSidebarOnTeamLeave", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) CountPostsAfter(channelId string, timestamp int64, userId string) (int, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.CountPostsAfter(channelId, timestamp, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.CountPostsAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) CreateDirectChannel(userId *model.User, otherUserId *model.User) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.CreateDirectChannel(userId, otherUserId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.CreateDirectChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) CreateInitialSidebarCategories(userId string, teamId string) error {
	start := timemodule.Now()

	err := s.ChannelStore.CreateInitialSidebarCategories(userId, teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.CreateInitialSidebarCategories", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) CreateSidebarCategory(userId string, teamId string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.CreateSidebarCategory(userId, teamId, newCategory)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.CreateSidebarCategory", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) Delete(channelId string, time int64) error {
	start := timemodule.Now()

	err := s.ChannelStore.Delete(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) DeleteSidebarCategory(categoryId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.DeleteSidebarCategory(categoryId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.DeleteSidebarCategory", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) DeleteSidebarChannelsByPreferences(preferences *model.Preferences) error {
	start := timemodule.Now()

	err := s.ChannelStore.DeleteSidebarChannelsByPreferences(preferences)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.DeleteSidebarChannelsByPreferences", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) Get(id string, allowFromCache bool) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.Get(id, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAll(teamId string) ([]*model.Channel, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAll(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllChannelMembersForUser(userId, allowFromCache, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllChannelMembersForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllChannelMembersNotifyPropsForChannel(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllChannelMembersNotifyPropsForChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllChannels(page int, perPage int, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllChannels(page, perPage, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllChannelsCount(opts store.ChannelSearchOpts) (int64, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllChannelsCount(opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllChannelsCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) ([]*model.ChannelForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllChannelsForExportAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllChannelsForExportAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) ([]*model.DirectChannelForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetAllDirectChannelsForExportAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetAllDirectChannelsForExportAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetByName(team_id string, name string, allowFromCache bool) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetByName(team_id, name, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetByNameIncludeDeleted(team_id, name, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetByNameIncludeDeleted", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) ([]*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetByNames(team_id, names, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetByNames", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelCounts(teamId string, userId string) (*model.ChannelCounts, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelCounts(teamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelCounts", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelMembersForExport(userId string, teamId string) ([]*model.ChannelMemberForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelMembersForExport(userId, teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelMembersForExport", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelMembersTimezones(channelId string) ([]model.StringMap, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelMembersTimezones(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelMembersTimezones", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelUnread(channelId string, userId string) (*model.ChannelUnread, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelUnread(channelId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelUnread", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannels(teamId string, userId string, includeDeleted bool, lastDeleteAt int) (*model.ChannelList, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannels(teamId, userId, includeDeleted, lastDeleteAt)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.Channel, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelsBatchForIndexing(startTime, endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelsBatchForIndexing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelsByIds(channelIds []string, includeDeleted bool) ([]*model.Channel, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelsByIds(channelIds, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelsByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) (model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetChannelsByScheme(schemeId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetChannelsByScheme", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetDeleted(team_id string, offset int, limit int, userId string) (*model.ChannelList, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetDeleted(team_id, offset, limit, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetDeleted", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetDeletedByName(team_id string, name string) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetDeletedByName(team_id, name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetDeletedByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetForPost(postId string) (*model.Channel, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetForPost(postId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetForPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetFromMaster(id string) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetFromMaster(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetFromMaster", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetGuestCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetGuestCount(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetGuestCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMember(channelId string, userId string) (*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMember(channelId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMemberCount(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMemberCountFromCache(channelId string) int64 {
	start := timemodule.Now()

	result := s.ChannelStore.GetMemberCountFromCache(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMemberCountFromCache", success, elapsed)
	}
	return result
}

func (s *TimerLayerChannelStore) GetMemberCountsByGroup(channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMemberCountsByGroup(channelID, includeTimezones)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMemberCountsByGroup", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMemberForPost(postId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMemberForPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMembers(channelId string, offset int, limit int) (*model.ChannelMembers, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMembers(channelId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMembersByIds(channelId string, userIds []string) (*model.ChannelMembers, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMembersByIds(channelId, userIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMembersByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMembersForUser(teamId string, userId string) (*model.ChannelMembers, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMembersForUser(teamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMembersForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMembersForUserWithPagination(teamId string, userId string, page int, perPage int) (*model.ChannelMembers, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMembersForUserWithPagination(teamId, userId, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMembersForUserWithPagination", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) (*model.ChannelList, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetMoreChannels(teamId, userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetMoreChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetPinnedPostCount(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetPinnedPostCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetPinnedPosts(channelId string) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetPinnedPosts(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetPinnedPosts", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetPrivateChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetPrivateChannelsForTeam(teamId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetPrivateChannelsForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetPublicChannelsByIdsForTeam(teamId, channelIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetPublicChannelsByIdsForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetPublicChannelsForTeam(teamId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetPublicChannelsForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetSidebarCategories(userId string, teamId string) (*model.OrderedSidebarCategories, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetSidebarCategories(userId, teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetSidebarCategories", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetSidebarCategory(categoryId string) (*model.SidebarCategoryWithChannels, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetSidebarCategory(categoryId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetSidebarCategory", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetSidebarCategoryOrder(userId string, teamId string) ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetSidebarCategoryOrder(userId, teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetSidebarCategoryOrder", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GetTeamChannels(teamId string) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GetTeamChannels(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GetTeamChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) GroupSyncedChannelCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.GroupSyncedChannelCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.GroupSyncedChannelCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) IncrementMentionCount(channelId string, userId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.IncrementMentionCount(channelId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.IncrementMentionCount", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateAllChannelMembersForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateAllChannelMembersForUser", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateCacheForChannelMembersNotifyProps(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateCacheForChannelMembersNotifyProps", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidateChannel(id string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateChannel(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateChannel", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidateChannelByName(teamId string, name string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateChannelByName(teamId, name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateChannelByName", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidateGuestCount(channelId string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateGuestCount(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateGuestCount", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidateMemberCount(channelId string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidateMemberCount(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidateMemberCount", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) InvalidatePinnedPostCount(channelId string) {
	start := timemodule.Now()

	s.ChannelStore.InvalidatePinnedPostCount(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.InvalidatePinnedPostCount", success, elapsed)
	}
}

func (s *TimerLayerChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	start := timemodule.Now()

	result := s.ChannelStore.IsUserInChannelUseCache(userId, channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.IsUserInChannelUseCache", success, elapsed)
	}
	return result
}

func (s *TimerLayerChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) (map[string]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.MigrateChannelMembers(fromChannelId, fromUserId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.MigrateChannelMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) MigratePublicChannels() error {
	start := timemodule.Now()

	err := s.ChannelStore.MigratePublicChannels()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.MigratePublicChannels", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) PermanentDelete(channelId string) error {
	start := timemodule.Now()

	err := s.ChannelStore.PermanentDelete(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.PermanentDelete", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) PermanentDeleteByTeam(teamId string) error {
	start := timemodule.Now()

	err := s.ChannelStore.PermanentDeleteByTeam(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.PermanentDeleteByTeam", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) PermanentDeleteMembersByChannel(channelId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.PermanentDeleteMembersByChannel(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.PermanentDeleteMembersByChannel", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) PermanentDeleteMembersByUser(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.PermanentDeleteMembersByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.PermanentDeleteMembersByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) RemoveAllDeactivatedMembers(channelId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.RemoveAllDeactivatedMembers(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.RemoveAllDeactivatedMembers", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) RemoveMember(channelId string, userId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.RemoveMember(channelId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.RemoveMember", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) RemoveMembers(channelId string, userIds []string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.RemoveMembers(channelId, userIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.RemoveMembers", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) ResetAllChannelSchemes() *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.ResetAllChannelSchemes()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.ResetAllChannelSchemes", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) Restore(channelId string, time int64) error {
	start := timemodule.Now()

	err := s.ChannelStore.Restore(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.Restore", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.Save(channel, maxChannelsPerTeam)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SaveDirectChannel(channel, member1, member2)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SaveDirectChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SaveMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SaveMember(member)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SaveMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SaveMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SaveMultipleMembers(members)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SaveMultipleMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SearchAllChannels(term string, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, int64, *model.AppError) {
	start := timemodule.Now()

	result, resultVar1, err := s.ChannelStore.SearchAllChannels(term, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchAllChannels", success, elapsed)
	}
	return result, resultVar1, err
}

func (s *TimerLayerChannelStore) SearchArchivedInTeam(teamId string, term string, userId string) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SearchArchivedInTeam(teamId, term, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchArchivedInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SearchForUserInTeam(userId string, teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SearchForUserInTeam(userId, teamId, term, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchForUserInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SearchGroupChannels(userId string, term string) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SearchGroupChannels(userId, term)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchGroupChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SearchInTeam(teamId, term, includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SearchMore(userId string, teamId string, term string) (*model.ChannelList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.SearchMore(userId, teamId, term)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SearchMore", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) error {
	start := timemodule.Now()

	err := s.ChannelStore.SetDeleteAt(channelId, deleteAt, updateAt)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.SetDeleteAt", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) Update(channel *model.Channel) (*model.Channel, error) {
	start := timemodule.Now()

	result, err := s.ChannelStore.Update(channel)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateLastViewedAt(channelIds []string, userId string) (map[string]int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UpdateLastViewedAt(channelIds, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateLastViewedAt", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateLastViewedAtPost(unreadPost *model.Post, userID string, mentionCount int) (*model.ChannelUnreadAt, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UpdateLastViewedAtPost(unreadPost, userID, mentionCount)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateLastViewedAtPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UpdateMember(member)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateMembersRole(channelID string, userIDs []string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.UpdateMembersRole(channelID, userIDs)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateMembersRole", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) UpdateMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UpdateMultipleMembers(members)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateMultipleMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateSidebarCategories(userId string, teamId string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UpdateSidebarCategories(userId, teamId, categories)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateSidebarCategories", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelStore) UpdateSidebarCategoryOrder(userId string, teamId string, categoryOrder []string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.UpdateSidebarCategoryOrder(userId, teamId, categoryOrder)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateSidebarCategoryOrder", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) UpdateSidebarChannelCategoryOnMove(channel *model.Channel, newTeamId string) *model.AppError {
	start := timemodule.Now()

	err := s.ChannelStore.UpdateSidebarChannelCategoryOnMove(channel, newTeamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateSidebarChannelCategoryOnMove", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) UpdateSidebarChannelsByPreferences(preferences *model.Preferences) error {
	start := timemodule.Now()

	err := s.ChannelStore.UpdateSidebarChannelsByPreferences(preferences)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UpdateSidebarChannelsByPreferences", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelStore) UserBelongsToChannels(userId string, channelIds []string) (bool, *model.AppError) {
	start := timemodule.Now()

	result, err := s.ChannelStore.UserBelongsToChannels(userId, channelIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelStore.UserBelongsToChannels", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) ([]*model.ChannelMemberHistoryResult, error) {
	start := timemodule.Now()

	result, err := s.ChannelMemberHistoryStore.GetUsersInChannelDuring(startTime, endTime, channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelMemberHistoryStore.GetUsersInChannelDuring", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) error {
	start := timemodule.Now()

	err := s.ChannelMemberHistoryStore.LogJoinEvent(userId, channelId, joinTime)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelMemberHistoryStore.LogJoinEvent", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) error {
	start := timemodule.Now()

	err := s.ChannelMemberHistoryStore.LogLeaveEvent(userId, channelId, leaveTime)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelMemberHistoryStore.LogLeaveEvent", success, elapsed)
	}
	return err
}

func (s *TimerLayerChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	start := timemodule.Now()

	result, err := s.ChannelMemberHistoryStore.PermanentDeleteBatch(endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ChannelMemberHistoryStore.PermanentDeleteBatch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerClusterDiscoveryStore) Cleanup() error {
	start := timemodule.Now()

	err := s.ClusterDiscoveryStore.Cleanup()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.Cleanup", success, elapsed)
	}
	return err
}

func (s *TimerLayerClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, error) {
	start := timemodule.Now()

	result, err := s.ClusterDiscoveryStore.Delete(discovery)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, error) {
	start := timemodule.Now()

	result, err := s.ClusterDiscoveryStore.Exists(discovery)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.Exists", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, error) {
	start := timemodule.Now()

	result, err := s.ClusterDiscoveryStore.GetAll(discoveryType, clusterName)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) error {
	start := timemodule.Now()

	err := s.ClusterDiscoveryStore.Save(discovery)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.Save", success, elapsed)
	}
	return err
}

func (s *TimerLayerClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) error {
	start := timemodule.Now()

	err := s.ClusterDiscoveryStore.SetLastPingAt(discovery)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ClusterDiscoveryStore.SetLastPingAt", success, elapsed)
	}
	return err
}

func (s *TimerLayerCommandStore) AnalyticsCommandCount(teamId string) (int64, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.AnalyticsCommandCount(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.AnalyticsCommandCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandStore) Delete(commandId string, time int64) error {
	start := timemodule.Now()

	err := s.CommandStore.Delete(commandId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerCommandStore) Get(id string) (*model.Command, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandStore) GetByTeam(teamId string) ([]*model.Command, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.GetByTeam(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.GetByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandStore) GetByTrigger(teamId string, trigger string) (*model.Command, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.GetByTrigger(teamId, trigger)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.GetByTrigger", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandStore) PermanentDeleteByTeam(teamId string) error {
	start := timemodule.Now()

	err := s.CommandStore.PermanentDeleteByTeam(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.PermanentDeleteByTeam", success, elapsed)
	}
	return err
}

func (s *TimerLayerCommandStore) PermanentDeleteByUser(userId string) error {
	start := timemodule.Now()

	err := s.CommandStore.PermanentDeleteByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.PermanentDeleteByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerCommandStore) Save(webhook *model.Command) (*model.Command, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.Save(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandStore) Update(hook *model.Command) (*model.Command, error) {
	start := timemodule.Now()

	result, err := s.CommandStore.Update(hook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandWebhookStore) Cleanup() {
	start := timemodule.Now()

	s.CommandWebhookStore.Cleanup()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandWebhookStore.Cleanup", success, elapsed)
	}
}

func (s *TimerLayerCommandWebhookStore) Get(id string) (*model.CommandWebhook, error) {
	start := timemodule.Now()

	result, err := s.CommandWebhookStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandWebhookStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandWebhookStore) Save(webhook *model.CommandWebhook) (*model.CommandWebhook, error) {
	start := timemodule.Now()

	result, err := s.CommandWebhookStore.Save(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandWebhookStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerCommandWebhookStore) TryUse(id string, limit int) error {
	start := timemodule.Now()

	err := s.CommandWebhookStore.TryUse(id, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("CommandWebhookStore.TryUse", success, elapsed)
	}
	return err
}

func (s *TimerLayerComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.ComplianceExport(compliance)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.ComplianceExport", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerComplianceStore) Get(id string) (*model.Compliance, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerComplianceStore) GetAll(offset int, limit int) (model.Compliances, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.GetAll(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.MessageExport(after, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.MessageExport", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.Save(compliance)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, error) {
	start := timemodule.Now()

	result, err := s.ComplianceStore.Update(compliance)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ComplianceStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) Delete(emoji *model.Emoji, time int64) error {
	start := timemodule.Now()

	err := s.EmojiStore.Delete(emoji, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerEmojiStore) Get(id string, allowFromCache bool) (*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.Get(id, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) GetByName(name string, allowFromCache bool) (*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.GetByName(name, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) GetList(offset int, limit int, sort string) ([]*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.GetList(offset, limit, sort)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.GetList", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) GetMultipleByName(names []string) ([]*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.GetMultipleByName(names)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.GetMultipleByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.Save(emoji)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerEmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {
	start := timemodule.Now()

	result, err := s.EmojiStore.Search(name, prefixOnly, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("EmojiStore.Search", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) AttachToPost(fileId string, postId string, creatorId string) error {
	start := timemodule.Now()

	err := s.FileInfoStore.AttachToPost(fileId, postId, creatorId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.AttachToPost", success, elapsed)
	}
	return err
}

func (s *TimerLayerFileInfoStore) ClearCaches() {
	start := timemodule.Now()

	s.FileInfoStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerFileInfoStore) DeleteForPost(postId string) (string, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.DeleteForPost(postId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.DeleteForPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) Get(id string) (*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) GetByPath(path string) (*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.GetByPath(path)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.GetByPath", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) GetForPost(postId string, readFromMaster bool, includeDeleted bool, allowFromCache bool) ([]*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.GetForPost(postId, readFromMaster, includeDeleted, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.GetForPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) GetForUser(userId string) ([]*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.GetForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.GetForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) GetWithOptions(page int, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.GetWithOptions(page, perPage, opt)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.GetWithOptions", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) InvalidateFileInfosForPostCache(postId string, deleted bool) {
	start := timemodule.Now()

	s.FileInfoStore.InvalidateFileInfosForPostCache(postId, deleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.InvalidateFileInfosForPostCache", success, elapsed)
	}
}

func (s *TimerLayerFileInfoStore) PermanentDelete(fileId string) error {
	start := timemodule.Now()

	err := s.FileInfoStore.PermanentDelete(fileId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.PermanentDelete", success, elapsed)
	}
	return err
}

func (s *TimerLayerFileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.PermanentDeleteBatch(endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.PermanentDeleteBatch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) PermanentDeleteByUser(userId string) (int64, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.PermanentDeleteByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.PermanentDeleteByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerFileInfoStore) Save(info *model.FileInfo) (*model.FileInfo, error) {
	start := timemodule.Now()

	result, err := s.FileInfoStore.Save(info)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("FileInfoStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) AdminRoleGroupsForSyncableMember(userID string, syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.AdminRoleGroupsForSyncableMember(userID, syncableID, syncableType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.AdminRoleGroupsForSyncableMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.ChannelMembersMinusGroupMembers(channelID, groupIDs, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.ChannelMembersMinusGroupMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) ChannelMembersToAdd(since int64, channelID *string) ([]*model.UserChannelIDPair, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.ChannelMembersToAdd(since, channelID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.ChannelMembersToAdd", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.ChannelMembersToRemove(channelID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.ChannelMembersToRemove", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.CountChannelMembersMinusGroupMembers(channelID, groupIDs)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.CountChannelMembersMinusGroupMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.CountGroupsByChannel(channelId, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.CountGroupsByChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.CountGroupsByTeam(teamId, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.CountGroupsByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.CountTeamMembersMinusGroupMembers(teamID, groupIDs)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.CountTeamMembersMinusGroupMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) Create(group *model.Group) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.Create(group)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.Create", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.CreateGroupSyncable(groupSyncable)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.CreateGroupSyncable", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) Delete(groupID string) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.Delete(groupID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.DeleteGroupSyncable(groupID, syncableID, syncableType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.DeleteGroupSyncable", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.DeleteMember(groupID, userID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.DeleteMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) DistinctGroupMemberCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.DistinctGroupMemberCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.DistinctGroupMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) Get(groupID string) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.Get(groupID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetAllBySource(groupSource)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetAllBySource", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetAllGroupSyncablesByGroupId", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetByIDs(groupIDs []string) ([]*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetByIDs(groupIDs)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetByIDs", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetByName(name, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetByRemoteID(remoteID, groupSource)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetByRemoteID", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetByUser(userId string) ([]*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetGroupSyncable(groupID, syncableID, syncableType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetGroupSyncable", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetGroups(page int, perPage int, opts model.GroupSearchOpts) ([]*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetGroups(page, perPage, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetGroups", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetGroupsAssociatedToChannelsByTeam(teamId, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetGroupsAssociatedToChannelsByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetGroupsByChannel(channelId, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetGroupsByChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetGroupsByTeam(teamId, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetGroupsByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetMemberCount(groupID string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetMemberCount(groupID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetMemberUsers(groupID string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetMemberUsers(groupID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetMemberUsers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetMemberUsersInTeam(groupID, teamID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetMemberUsersInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetMemberUsersNotInChannel(groupID, channelID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetMemberUsersNotInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GetMemberUsersPage(groupID, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GetMemberUsersPage", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GroupChannelCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GroupChannelCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GroupChannelCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GroupCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GroupCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GroupCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GroupCountWithAllowReference() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GroupCountWithAllowReference()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GroupCountWithAllowReference", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GroupMemberCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GroupMemberCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GroupMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) GroupTeamCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.GroupTeamCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.GroupTeamCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) PermanentDeleteMembersByUser(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.GroupStore.PermanentDeleteMembersByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.PermanentDeleteMembersByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerGroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.PermittedSyncableAdmins(syncableID, syncableType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.PermittedSyncableAdmins", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.TeamMembersMinusGroupMembers(teamID, groupIDs, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.TeamMembersMinusGroupMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) TeamMembersToAdd(since int64, teamID *string) ([]*model.UserTeamIDPair, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.TeamMembersToAdd(since, teamID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.TeamMembersToAdd", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.TeamMembersToRemove(teamID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.TeamMembersToRemove", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) Update(group *model.Group) (*model.Group, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.Update(group)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.UpdateGroupSyncable(groupSyncable)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.UpdateGroupSyncable", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerGroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.GroupStore.UpsertMember(groupID, userID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("GroupStore.UpsertMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) Delete(id string) (string, error) {
	start := timemodule.Now()

	result, err := s.JobStore.Delete(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) Get(id string) (*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetAllByStatus(status string) ([]*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetAllByStatus(status)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetAllByStatus", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetAllByType(jobType string) ([]*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetAllByType(jobType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetAllByType", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetAllByTypePage(jobType string, offset int, limit int) ([]*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetAllByTypePage(jobType, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetAllByTypePage", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetAllPage(offset int, limit int) ([]*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetAllPage(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetAllPage", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetCountByStatusAndType(status string, jobType string) (int64, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetCountByStatusAndType(status, jobType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetCountByStatusAndType", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetNewestJobByStatusAndType(status string, jobType string) (*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetNewestJobByStatusAndType(status, jobType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetNewestJobByStatusAndType", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) GetNewestJobByStatusesAndType(statuses []string, jobType string) (*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.GetNewestJobByStatusesAndType(statuses, jobType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.GetNewestJobByStatusesAndType", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) Save(job *model.Job) (*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.Save(job)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) UpdateOptimistically(job *model.Job, currentStatus string) (bool, error) {
	start := timemodule.Now()

	result, err := s.JobStore.UpdateOptimistically(job, currentStatus)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.UpdateOptimistically", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) UpdateStatus(id string, status string) (*model.Job, error) {
	start := timemodule.Now()

	result, err := s.JobStore.UpdateStatus(id, status)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.UpdateStatus", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerJobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) (bool, error) {
	start := timemodule.Now()

	result, err := s.JobStore.UpdateStatusOptimistically(id, currentStatus, newStatus)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("JobStore.UpdateStatusOptimistically", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerLicenseStore) Get(id string) (*model.LicenseRecord, error) {
	start := timemodule.Now()

	result, err := s.LicenseStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("LicenseStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, error) {
	start := timemodule.Now()

	result, err := s.LicenseStore.Save(license)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("LicenseStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerLinkMetadataStore) Get(url string, timestamp int64) (*model.LinkMetadata, error) {
	start := timemodule.Now()

	result, err := s.LinkMetadataStore.Get(url, timestamp)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("LinkMetadataStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerLinkMetadataStore) Save(linkMetadata *model.LinkMetadata) (*model.LinkMetadata, error) {
	start := timemodule.Now()

	result, err := s.LinkMetadataStore.Save(linkMetadata)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("LinkMetadataStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) DeleteApp(id string) error {
	start := timemodule.Now()

	err := s.OAuthStore.DeleteApp(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.DeleteApp", success, elapsed)
	}
	return err
}

func (s *TimerLayerOAuthStore) GetAccessData(token string) (*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAccessData(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAccessData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetAccessDataByRefreshToken(token string) (*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAccessDataByRefreshToken(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAccessDataByRefreshToken", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetAccessDataByUserForApp(userId string, clientId string) ([]*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAccessDataByUserForApp(userId, clientId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAccessDataByUserForApp", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetApp(id string) (*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetApp(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetApp", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetAppByUser(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAppByUser(userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAppByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetApps(offset int, limit int) ([]*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetApps(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetApps", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetAuthData(code string) (*model.AuthData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAuthData(code)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAuthData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetAuthorizedApps(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetAuthorizedApps(userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetAuthorizedApps", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) GetPreviousAccessData(userId string, clientId string) (*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.GetPreviousAccessData(userId, clientId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.GetPreviousAccessData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) PermanentDeleteAuthDataByUser(userId string) error {
	start := timemodule.Now()

	err := s.OAuthStore.PermanentDeleteAuthDataByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.PermanentDeleteAuthDataByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerOAuthStore) RemoveAccessData(token string) error {
	start := timemodule.Now()

	err := s.OAuthStore.RemoveAccessData(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.RemoveAccessData", success, elapsed)
	}
	return err
}

func (s *TimerLayerOAuthStore) RemoveAllAccessData() error {
	start := timemodule.Now()

	err := s.OAuthStore.RemoveAllAccessData()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.RemoveAllAccessData", success, elapsed)
	}
	return err
}

func (s *TimerLayerOAuthStore) RemoveAuthData(code string) error {
	start := timemodule.Now()

	err := s.OAuthStore.RemoveAuthData(code)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.RemoveAuthData", success, elapsed)
	}
	return err
}

func (s *TimerLayerOAuthStore) SaveAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.SaveAccessData(accessData)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.SaveAccessData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) SaveApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.SaveApp(app)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.SaveApp", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) SaveAuthData(authData *model.AuthData) (*model.AuthData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.SaveAuthData(authData)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.SaveAuthData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) UpdateAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.UpdateAccessData(accessData)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.UpdateAccessData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerOAuthStore) UpdateApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	start := timemodule.Now()

	result, err := s.OAuthStore.UpdateApp(app)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("OAuthStore.UpdateApp", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) CompareAndDelete(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.CompareAndDelete(keyVal, oldValue)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.CompareAndDelete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.CompareAndSet(keyVal, oldValue)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.CompareAndSet", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) Delete(pluginId string, key string) error {
	start := timemodule.Now()

	err := s.PluginStore.Delete(pluginId, key)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerPluginStore) DeleteAllExpired() error {
	start := timemodule.Now()

	err := s.PluginStore.DeleteAllExpired()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.DeleteAllExpired", success, elapsed)
	}
	return err
}

func (s *TimerLayerPluginStore) DeleteAllForPlugin(PluginId string) error {
	start := timemodule.Now()

	err := s.PluginStore.DeleteAllForPlugin(PluginId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.DeleteAllForPlugin", success, elapsed)
	}
	return err
}

func (s *TimerLayerPluginStore) Get(pluginId string, key string) (*model.PluginKeyValue, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.Get(pluginId, key)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) List(pluginId string, page int, perPage int) ([]string, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.List(pluginId, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.List", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) (*model.PluginKeyValue, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.SaveOrUpdate(keyVal)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.SaveOrUpdate", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPluginStore) SetWithOptions(pluginId string, key string, value []byte, options model.PluginKVSetOptions) (bool, error) {
	start := timemodule.Now()

	result, err := s.PluginStore.SetWithOptions(pluginId, key, value, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PluginStore.SetWithOptions", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.AnalyticsPostCount(teamId, mustHaveFile, mustHaveHashtag)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.AnalyticsPostCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.AnalyticsPostCountsByDay(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.AnalyticsPostCountsByDay", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) (model.AnalyticsRows, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.AnalyticsUserCountsWithPostsByDay(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.AnalyticsUserCountsWithPostsByDay", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) ClearCaches() {
	start := timemodule.Now()

	s.PostStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerPostStore) Delete(postId string, time int64, deleteByID string) error {
	start := timemodule.Now()

	err := s.PostStore.Delete(postId, time, deleteByID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerPostStore) Get(id string, skipFetchThreads bool) (*model.PostList, error) {
	start := timemodule.Now()

	result, err := s.PostStore.Get(id, skipFetchThreads)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) ([]*model.DirectPostForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetDirectPostParentsForExportAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetDirectPostParentsForExportAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetEtag(channelId string, allowFromCache bool) string {
	start := timemodule.Now()

	result := s.PostStore.GetEtag(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetEtag", success, elapsed)
	}
	return result
}

func (s *TimerLayerPostStore) GetFlaggedPosts(userId string, offset int, limit int) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetFlaggedPosts(userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetFlaggedPosts", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetFlaggedPostsForChannel(userId, channelId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetFlaggedPostsForChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetFlaggedPostsForTeam(userId, teamId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetFlaggedPostsForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetMaxPostSize() int {
	start := timemodule.Now()

	result := s.PostStore.GetMaxPostSize()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetMaxPostSize", success, elapsed)
	}
	return result
}

func (s *TimerLayerPostStore) GetOldest() (*model.Post, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetOldest()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetOldest", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetOldestEntityCreationTime() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetOldestEntityCreationTime()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetOldestEntityCreationTime", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetParentsForExportAfter(limit int, afterId string) ([]*model.PostForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetParentsForExportAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetParentsForExportAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostAfterTime(channelId string, time int64) (*model.Post, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostAfterTime(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostAfterTime", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostIdAfterTime(channelId string, time int64) (string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostIdAfterTime(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostIdAfterTime", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostIdBeforeTime(channelId string, time int64) (string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostIdBeforeTime(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostIdBeforeTime", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPosts(options model.GetPostsOptions, allowFromCache bool) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPosts(options, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPosts", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsAfter(options model.GetPostsOptions) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsAfter(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.PostForIndexing, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsBatchForIndexing(startTime, endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsBatchForIndexing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsBefore(options model.GetPostsOptions) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsBefore(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsBefore", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsByIds(postIds []string) ([]*model.Post, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsByIds(postIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsCreatedAt(channelId string, time int64) ([]*model.Post, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsCreatedAt(channelId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsCreatedAt", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetPostsSince(options, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetPostsSince", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetRepliesForExport(parentId string) ([]*model.ReplyForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.GetRepliesForExport(parentId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetRepliesForExport", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) GetSingle(id string) (*model.Post, error) {
	start := timemodule.Now()

	result, err := s.PostStore.GetSingle(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.GetSingle", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) InvalidateLastPostTimeCache(channelId string) {
	start := timemodule.Now()

	s.PostStore.InvalidateLastPostTimeCache(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.InvalidateLastPostTimeCache", success, elapsed)
	}
}

func (s *TimerLayerPostStore) Overwrite(post *model.Post) (*model.Post, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.Overwrite(post)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Overwrite", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, *model.AppError) {
	start := timemodule.Now()

	result, resultVar1, err := s.PostStore.OverwriteMultiple(posts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.OverwriteMultiple", success, elapsed)
	}
	return result, resultVar1, err
}

func (s *TimerLayerPostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.PermanentDeleteBatch(endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.PermanentDeleteBatch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) PermanentDeleteByChannel(channelId string) error {
	start := timemodule.Now()

	err := s.PostStore.PermanentDeleteByChannel(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.PermanentDeleteByChannel", success, elapsed)
	}
	return err
}

func (s *TimerLayerPostStore) PermanentDeleteByUser(userId string) error {
	start := timemodule.Now()

	err := s.PostStore.PermanentDeleteByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.PermanentDeleteByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerPostStore) Save(post *model.Post) (*model.Post, error) {
	start := timemodule.Now()

	result, err := s.PostStore.Save(post)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	start := timemodule.Now()

	result, resultVar1, err := s.PostStore.SaveMultiple(posts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.SaveMultiple", success, elapsed)
	}
	return result, resultVar1, err
}

func (s *TimerLayerPostStore) Search(teamId string, userId string, params *model.SearchParams) (*model.PostList, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.Search(teamId, userId, params)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Search", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) SearchPostsInTeamForUser(paramsList []*model.SearchParams, userId string, teamId string, isOrSearch bool, includeDeletedChannels bool, page int, perPage int) (*model.PostSearchResults, *model.AppError) {
	start := timemodule.Now()

	result, err := s.PostStore.SearchPostsInTeamForUser(paramsList, userId, teamId, isOrSearch, includeDeletedChannels, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.SearchPostsInTeamForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPostStore) Update(newPost *model.Post, oldPost *model.Post) (*model.Post, error) {
	start := timemodule.Now()

	result, err := s.PostStore.Update(newPost, oldPost)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PostStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {
	start := timemodule.Now()

	result, err := s.PreferenceStore.CleanupFlagsBatch(limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.CleanupFlagsBatch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPreferenceStore) Delete(userId string, category string, name string) error {
	start := timemodule.Now()

	err := s.PreferenceStore.Delete(userId, category, name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerPreferenceStore) DeleteCategory(userId string, category string) error {
	start := timemodule.Now()

	err := s.PreferenceStore.DeleteCategory(userId, category)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.DeleteCategory", success, elapsed)
	}
	return err
}

func (s *TimerLayerPreferenceStore) DeleteCategoryAndName(category string, name string) error {
	start := timemodule.Now()

	err := s.PreferenceStore.DeleteCategoryAndName(category, name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.DeleteCategoryAndName", success, elapsed)
	}
	return err
}

func (s *TimerLayerPreferenceStore) Get(userId string, category string, name string) (*model.Preference, error) {
	start := timemodule.Now()

	result, err := s.PreferenceStore.Get(userId, category, name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPreferenceStore) GetAll(userId string) (model.Preferences, error) {
	start := timemodule.Now()

	result, err := s.PreferenceStore.GetAll(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPreferenceStore) GetCategory(userId string, category string) (model.Preferences, error) {
	start := timemodule.Now()

	result, err := s.PreferenceStore.GetCategory(userId, category)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.GetCategory", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerPreferenceStore) PermanentDeleteByUser(userId string) error {
	start := timemodule.Now()

	err := s.PreferenceStore.PermanentDeleteByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.PermanentDeleteByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerPreferenceStore) Save(preferences *model.Preferences) error {
	start := timemodule.Now()

	err := s.PreferenceStore.Save(preferences)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("PreferenceStore.Save", success, elapsed)
	}
	return err
}

func (s *TimerLayerReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, error) {
	start := timemodule.Now()

	result, err := s.ReactionStore.BulkGetForPosts(postIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.BulkGetForPosts", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, error) {
	start := timemodule.Now()

	result, err := s.ReactionStore.Delete(reaction)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerReactionStore) DeleteAllWithEmojiName(emojiName string) error {
	start := timemodule.Now()

	err := s.ReactionStore.DeleteAllWithEmojiName(emojiName)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.DeleteAllWithEmojiName", success, elapsed)
	}
	return err
}

func (s *TimerLayerReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, error) {
	start := timemodule.Now()

	result, err := s.ReactionStore.GetForPost(postId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.GetForPost", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	start := timemodule.Now()

	result, err := s.ReactionStore.PermanentDeleteBatch(endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.PermanentDeleteBatch", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerReactionStore) Save(reaction *model.Reaction) (*model.Reaction, error) {
	start := timemodule.Now()

	result, err := s.ReactionStore.Save(reaction)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("ReactionStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) AllChannelSchemeRoles() ([]*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.AllChannelSchemeRoles()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.AllChannelSchemeRoles", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) ChannelHigherScopedPermissions(roleNames []string) (map[string]*model.RolePermissions, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.ChannelHigherScopedPermissions(roleNames)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.ChannelHigherScopedPermissions", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) ChannelRolesUnderTeamRole(roleName string) ([]*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.ChannelRolesUnderTeamRole(roleName)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.ChannelRolesUnderTeamRole", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) Delete(roleId string) (*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.Delete(roleId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) Get(roleId string) (*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.Get(roleId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) GetAll() ([]*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.GetAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) GetByName(name string) (*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.GetByName(name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) GetByNames(names []string) ([]*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.GetByNames(names)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.GetByNames", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerRoleStore) PermanentDeleteAll() error {
	start := timemodule.Now()

	err := s.RoleStore.PermanentDeleteAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.PermanentDeleteAll", success, elapsed)
	}
	return err
}

func (s *TimerLayerRoleStore) Save(role *model.Role) (*model.Role, error) {
	start := timemodule.Now()

	result, err := s.RoleStore.Save(role)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("RoleStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) CountByScope(scope string) (int64, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.CountByScope(scope)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.CountByScope", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) CountWithoutPermission(scope string, permissionID string, roleScope model.RoleScope, roleType model.RoleType) (int64, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.CountWithoutPermission(scope, permissionID, roleScope, roleType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.CountWithoutPermission", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) Delete(schemeId string) (*model.Scheme, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.Delete(schemeId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.Delete", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) Get(schemeId string) (*model.Scheme, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.Get(schemeId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) GetAllPage(scope string, offset int, limit int) ([]*model.Scheme, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.GetAllPage(scope, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.GetAllPage", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) GetByName(schemeName string) (*model.Scheme, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.GetByName(schemeName)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSchemeStore) PermanentDeleteAll() error {
	start := timemodule.Now()

	err := s.SchemeStore.PermanentDeleteAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.PermanentDeleteAll", success, elapsed)
	}
	return err
}

func (s *TimerLayerSchemeStore) Save(scheme *model.Scheme) (*model.Scheme, error) {
	start := timemodule.Now()

	result, err := s.SchemeStore.Save(scheme)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SchemeStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) AnalyticsSessionCount() (int64, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.AnalyticsSessionCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.AnalyticsSessionCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) Cleanup(expiryTime int64, batchSize int64) {
	start := timemodule.Now()

	s.SessionStore.Cleanup(expiryTime, batchSize)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.Cleanup", success, elapsed)
	}
}

func (s *TimerLayerSessionStore) Get(sessionIdOrToken string) (*model.Session, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.Get(sessionIdOrToken)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) GetSessions(userId string) ([]*model.Session, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.GetSessions(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.GetSessions", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) GetSessionsExpired(thresholdMillis int64, mobileOnly bool, unnotifiedOnly bool) ([]*model.Session, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.GetSessionsExpired(thresholdMillis, mobileOnly, unnotifiedOnly)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.GetSessionsExpired", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) GetSessionsWithActiveDeviceIds(userId string) ([]*model.Session, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.GetSessionsWithActiveDeviceIds(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.GetSessionsWithActiveDeviceIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) PermanentDeleteSessionsByUser(teamId string) error {
	start := timemodule.Now()

	err := s.SessionStore.PermanentDeleteSessionsByUser(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.PermanentDeleteSessionsByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) Remove(sessionIdOrToken string) error {
	start := timemodule.Now()

	err := s.SessionStore.Remove(sessionIdOrToken)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.Remove", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) RemoveAllSessions() error {
	start := timemodule.Now()

	err := s.SessionStore.RemoveAllSessions()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.RemoveAllSessions", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) Save(session *model.Session) (*model.Session, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.Save(session)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) (string, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.UpdateDeviceId(id, deviceId, expiresAt)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateDeviceId", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSessionStore) UpdateExpiredNotify(sessionid string, notified bool) error {
	start := timemodule.Now()

	err := s.SessionStore.UpdateExpiredNotify(sessionid, notified)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateExpiredNotify", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) UpdateExpiresAt(sessionId string, time int64) error {
	start := timemodule.Now()

	err := s.SessionStore.UpdateExpiresAt(sessionId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateExpiresAt", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) UpdateLastActivityAt(sessionId string, time int64) error {
	start := timemodule.Now()

	err := s.SessionStore.UpdateLastActivityAt(sessionId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateLastActivityAt", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) UpdateProps(session *model.Session) error {
	start := timemodule.Now()

	err := s.SessionStore.UpdateProps(session)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateProps", success, elapsed)
	}
	return err
}

func (s *TimerLayerSessionStore) UpdateRoles(userId string, roles string) (string, error) {
	start := timemodule.Now()

	result, err := s.SessionStore.UpdateRoles(userId, roles)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SessionStore.UpdateRoles", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerStatusStore) Get(userId string) (*model.Status, error) {
	start := timemodule.Now()

	result, err := s.StatusStore.Get(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerStatusStore) GetByIds(userIds []string) ([]*model.Status, error) {
	start := timemodule.Now()

	result, err := s.StatusStore.GetByIds(userIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.GetByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerStatusStore) GetTotalActiveUsersCount() (int64, error) {
	start := timemodule.Now()

	result, err := s.StatusStore.GetTotalActiveUsersCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.GetTotalActiveUsersCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerStatusStore) ResetAll() error {
	start := timemodule.Now()

	err := s.StatusStore.ResetAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.ResetAll", success, elapsed)
	}
	return err
}

func (s *TimerLayerStatusStore) SaveOrUpdate(status *model.Status) error {
	start := timemodule.Now()

	err := s.StatusStore.SaveOrUpdate(status)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.SaveOrUpdate", success, elapsed)
	}
	return err
}

func (s *TimerLayerStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) error {
	start := timemodule.Now()

	err := s.StatusStore.UpdateLastActivityAt(userId, lastActivityAt)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("StatusStore.UpdateLastActivityAt", success, elapsed)
	}
	return err
}

func (s *TimerLayerSystemStore) Get() (model.StringMap, error) {
	start := timemodule.Now()

	result, err := s.SystemStore.Get()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSystemStore) GetByName(name string) (*model.System, error) {
	start := timemodule.Now()

	result, err := s.SystemStore.GetByName(name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSystemStore) InsertIfExists(system *model.System) (*model.System, error) {
	start := timemodule.Now()

	result, err := s.SystemStore.InsertIfExists(system)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.InsertIfExists", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSystemStore) PermanentDeleteByName(name string) (*model.System, error) {
	start := timemodule.Now()

	result, err := s.SystemStore.PermanentDeleteByName(name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.PermanentDeleteByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerSystemStore) Save(system *model.System) error {
	start := timemodule.Now()

	err := s.SystemStore.Save(system)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.Save", success, elapsed)
	}
	return err
}

func (s *TimerLayerSystemStore) SaveOrUpdate(system *model.System) error {
	start := timemodule.Now()

	err := s.SystemStore.SaveOrUpdate(system)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.SaveOrUpdate", success, elapsed)
	}
	return err
}

func (s *TimerLayerSystemStore) Update(system *model.System) error {
	start := timemodule.Now()

	err := s.SystemStore.Update(system)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("SystemStore.Update", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.AnalyticsGetTeamCountForScheme(schemeId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.AnalyticsGetTeamCountForScheme", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) AnalyticsPrivateTeamCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.AnalyticsPrivateTeamCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.AnalyticsPrivateTeamCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) AnalyticsPublicTeamCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.AnalyticsPublicTeamCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.AnalyticsPublicTeamCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) AnalyticsTeamCount(includeDeleted bool) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.AnalyticsTeamCount(includeDeleted)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.AnalyticsTeamCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) ClearAllCustomRoleAssignments() *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.ClearAllCustomRoleAssignments()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.ClearAllCustomRoleAssignments", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) ClearCaches() {
	start := timemodule.Now()

	s.TeamStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerTeamStore) Get(id string) (*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetActiveMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetActiveMemberCount(teamId, restrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetActiveMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAll() ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllForExportAfter(limit int, afterId string) ([]*model.TeamForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllForExportAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllForExportAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllPage(offset int, limit int) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllPage(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllPage", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllPrivateTeamListing() ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllPrivateTeamListing()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllPrivateTeamListing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllPrivateTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllPrivateTeamPageListing(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllPrivateTeamPageListing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllPublicTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllPublicTeamPageListing(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllPublicTeamPageListing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllTeamListing() ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllTeamListing()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllTeamListing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetAllTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetAllTeamPageListing(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetAllTeamPageListing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetByInviteId(inviteId string) (*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetByInviteId(inviteId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetByInviteId", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetByName(name string) (*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetByName(name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetByName", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetByNames(name []string) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetByNames(name)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetByNames", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetChannelUnreadsForAllTeams(excludeTeamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetChannelUnreadsForAllTeams", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetChannelUnreadsForTeam(teamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetChannelUnreadsForTeam(teamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetChannelUnreadsForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetMember(teamId string, userId string) (*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetMember(teamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetMembers(teamId string, offset int, limit int, teamMembersGetOptions *model.TeamMembersGetOptions) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetMembers(teamId, offset, limit, teamMembersGetOptions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetMembersByIds(teamId, userIds, restrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetMembersByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTeamMembersForExport(userId string) ([]*model.TeamMemberForExport, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTeamMembersForExport(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTeamMembersForExport", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTeamsByScheme(schemeId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTeamsByScheme", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTeamsByUserId(userId string) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTeamsByUserId(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTeamsByUserId", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTeamsForUser(userId string) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTeamsForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTeamsForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTeamsForUserWithPagination(userId, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTeamsForUserWithPagination", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetTotalMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetTotalMemberCount(teamId, restrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetTotalMemberCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GetUserTeamIds(userId string, allowFromCache bool) ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GetUserTeamIds(userId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GetUserTeamIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) GroupSyncedTeamCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.GroupSyncedTeamCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.GroupSyncedTeamCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) InvalidateAllTeamIdsForUser(userId string) {
	start := timemodule.Now()

	s.TeamStore.InvalidateAllTeamIdsForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.InvalidateAllTeamIdsForUser", success, elapsed)
	}
}

func (s *TimerLayerTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) (map[string]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.MigrateTeamMembers(fromTeamId, fromUserId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.MigrateTeamMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) PermanentDelete(teamId string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.PermanentDelete(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.PermanentDelete", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) RemoveAllMembersByTeam(teamId string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.RemoveAllMembersByTeam(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.RemoveAllMembersByTeam", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) RemoveAllMembersByUser(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.RemoveAllMembersByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.RemoveAllMembersByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) RemoveMember(teamId string, userId string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.RemoveMember(teamId, userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.RemoveMember", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) RemoveMembers(teamId string, userIds []string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.RemoveMembers(teamId, userIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.RemoveMembers", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) ResetAllTeamSchemes() *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.ResetAllTeamSchemes()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.ResetAllTeamSchemes", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) Save(team *model.Team) (*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.Save(team)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) (*model.TeamMember, error) {
	start := timemodule.Now()

	result, err := s.TeamStore.SaveMember(member, maxUsersPerTeam)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SaveMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) SaveMultipleMembers(members []*model.TeamMember, maxUsersPerTeam int) ([]*model.TeamMember, error) {
	start := timemodule.Now()

	result, err := s.TeamStore.SaveMultipleMembers(members, maxUsersPerTeam)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SaveMultipleMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) SearchAll(term string, opts *model.TeamSearch) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.SearchAll(term, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SearchAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) SearchAllPaged(term string, opts *model.TeamSearch) ([]*model.Team, int64, *model.AppError) {
	start := timemodule.Now()

	result, resultVar1, err := s.TeamStore.SearchAllPaged(term, opts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SearchAllPaged", success, elapsed)
	}
	return result, resultVar1, err
}

func (s *TimerLayerTeamStore) SearchOpen(term string) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.SearchOpen(term)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SearchOpen", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) SearchPrivate(term string) ([]*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.SearchPrivate(term)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.SearchPrivate", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.Update(team)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.UpdateLastTeamIconUpdate(teamId, curTime)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.UpdateLastTeamIconUpdate", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) UpdateMember(member *model.TeamMember) (*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.UpdateMember(member)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.UpdateMember", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) UpdateMembersRole(teamID string, userIDs []string) *model.AppError {
	start := timemodule.Now()

	err := s.TeamStore.UpdateMembersRole(teamID, userIDs)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.UpdateMembersRole", success, elapsed)
	}
	return err
}

func (s *TimerLayerTeamStore) UpdateMultipleMembers(members []*model.TeamMember) ([]*model.TeamMember, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.UpdateMultipleMembers(members)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.UpdateMultipleMembers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTeamStore) UserBelongsToTeams(userId string, teamIds []string) (bool, *model.AppError) {
	start := timemodule.Now()

	result, err := s.TeamStore.UserBelongsToTeams(userId, teamIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TeamStore.UserBelongsToTeams", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTermsOfServiceStore) Get(id string, allowFromCache bool) (*model.TermsOfService, error) {
	start := timemodule.Now()

	result, err := s.TermsOfServiceStore.Get(id, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TermsOfServiceStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTermsOfServiceStore) GetLatest(allowFromCache bool) (*model.TermsOfService, error) {
	start := timemodule.Now()

	result, err := s.TermsOfServiceStore.GetLatest(allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TermsOfServiceStore.GetLatest", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTermsOfServiceStore) Save(termsOfService *model.TermsOfService) (*model.TermsOfService, error) {
	start := timemodule.Now()

	result, err := s.TermsOfServiceStore.Save(termsOfService)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TermsOfServiceStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTokenStore) Cleanup() {
	start := timemodule.Now()

	s.TokenStore.Cleanup()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TokenStore.Cleanup", success, elapsed)
	}
}

func (s *TimerLayerTokenStore) Delete(token string) error {
	start := timemodule.Now()

	err := s.TokenStore.Delete(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TokenStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerTokenStore) GetByToken(token string) (*model.Token, error) {
	start := timemodule.Now()

	result, err := s.TokenStore.GetByToken(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TokenStore.GetByToken", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerTokenStore) RemoveAllTokensByType(tokenType string) error {
	start := timemodule.Now()

	err := s.TokenStore.RemoveAllTokensByType(tokenType)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TokenStore.RemoveAllTokensByType", success, elapsed)
	}
	return err
}

func (s *TimerLayerTokenStore) Save(recovery *model.Token) error {
	start := timemodule.Now()

	err := s.TokenStore.Save(recovery)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("TokenStore.Save", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) AnalyticsActiveCount(time int64, options model.UserCountOptions) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.AnalyticsActiveCount(time, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.AnalyticsActiveCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) AnalyticsGetGuestCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.AnalyticsGetGuestCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.AnalyticsGetGuestCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) AnalyticsGetInactiveUsersCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.AnalyticsGetInactiveUsersCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.AnalyticsGetInactiveUsersCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) AnalyticsGetSystemAdminCount() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.AnalyticsGetSystemAdminCount()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.AnalyticsGetSystemAdminCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) AutocompleteUsersInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) (*model.UserAutocompleteInChannel, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.AutocompleteUsersInChannel(teamId, channelId, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.AutocompleteUsersInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) ClearAllCustomRoleAssignments() *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.ClearAllCustomRoleAssignments()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.ClearAllCustomRoleAssignments", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) ClearCaches() {
	start := timemodule.Now()

	s.UserStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerUserStore) Count(options model.UserCountOptions) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.Count(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.Count", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) DeactivateGuests() ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.DeactivateGuests()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.DeactivateGuests", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) DemoteUserToGuest(userID string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.DemoteUserToGuest(userID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.DemoteUserToGuest", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) Get(id string) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.Get(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAll() ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAll()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAllAfter(limit int, afterId string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAllAfter(limit, afterId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAllAfter", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAllNotInAuthService(authServices []string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAllNotInAuthService(authServices)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAllNotInAuthService", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAllProfiles(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAllProfiles", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) (map[string]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAllProfilesInChannel(channelId, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAllProfilesInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAllUsingAuthService(authService string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAllUsingAuthService(authService)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAllUsingAuthService", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetAnyUnreadPostCountForChannel(userId, channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetAnyUnreadPostCountForChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetByAuth(authData *string, authService string) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetByAuth(authData, authService)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetByAuth", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetByEmail(email string) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetByEmail(email)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetByEmail", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetByUsername(username string) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetByUsername(username)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetByUsername", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetChannelGroupUsers(channelID string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetChannelGroupUsers(channelID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetChannelGroupUsers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetEtagForAllProfiles() string {
	start := timemodule.Now()

	result := s.UserStore.GetEtagForAllProfiles()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetEtagForAllProfiles", success, elapsed)
	}
	return result
}

func (s *TimerLayerUserStore) GetEtagForProfiles(teamId string) string {
	start := timemodule.Now()

	result := s.UserStore.GetEtagForProfiles(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetEtagForProfiles", success, elapsed)
	}
	return result
}

func (s *TimerLayerUserStore) GetEtagForProfilesNotInTeam(teamId string) string {
	start := timemodule.Now()

	result := s.UserStore.GetEtagForProfilesNotInTeam(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetEtagForProfilesNotInTeam", success, elapsed)
	}
	return result
}

func (s *TimerLayerUserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetForLogin(loginId, allowSignInWithUsername, allowSignInWithEmail)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetForLogin", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetKnownUsers(userID string) ([]string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetKnownUsers(userID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetKnownUsers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetNewUsersForTeam(teamId, offset, limit, viewRestrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetNewUsersForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfileByGroupChannelIdsForUser(userId string, channelIds []string) (map[string][]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfileByGroupChannelIdsForUser(userId, channelIds)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfileByGroupChannelIdsForUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfileByIds(userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfileByIds(userIds, options, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfileByIds", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfiles(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfiles", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesByUsernames(usernames, viewRestrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesByUsernames", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesInChannel(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesInChannel(channelId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesInChannelByStatus(channelId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesInChannelByStatus", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesNotInChannel(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesNotInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesNotInTeam(teamId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesNotInTeam(teamId, groupConstrained, offset, limit, viewRestrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesNotInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetProfilesWithoutTeam(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetProfilesWithoutTeam(options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetProfilesWithoutTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetRecentlyActiveUsersForTeam(teamId, offset, limit, viewRestrictions)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetRecentlyActiveUsersForTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetSystemAdminProfiles() (map[string]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetSystemAdminProfiles()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetSystemAdminProfiles", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetTeamGroupUsers(teamID string) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetTeamGroupUsers(teamID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetTeamGroupUsers", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetUnreadCount(userId string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetUnreadCount(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetUnreadCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetUnreadCountForChannel(userId string, channelId string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetUnreadCountForChannel(userId, channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetUnreadCountForChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.UserForIndexing, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.GetUsersBatchForIndexing(startTime, endTime, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.GetUsersBatchForIndexing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) InferSystemInstallDate() (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.InferSystemInstallDate()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.InferSystemInstallDate", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) InvalidateProfileCacheForUser(userId string) {
	start := timemodule.Now()

	s.UserStore.InvalidateProfileCacheForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.InvalidateProfileCacheForUser", success, elapsed)
	}
}

func (s *TimerLayerUserStore) InvalidateProfilesInChannelCache(channelId string) {
	start := timemodule.Now()

	s.UserStore.InvalidateProfilesInChannelCache(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.InvalidateProfilesInChannelCache", success, elapsed)
	}
}

func (s *TimerLayerUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {
	start := timemodule.Now()

	s.UserStore.InvalidateProfilesInChannelCacheByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.InvalidateProfilesInChannelCacheByUser", success, elapsed)
	}
}

func (s *TimerLayerUserStore) PermanentDelete(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.PermanentDelete(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.PermanentDelete", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) PromoteGuestToUser(userID string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.PromoteGuestToUser(userID)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.PromoteGuestToUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) ResetLastPictureUpdate(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.ResetLastPictureUpdate(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.ResetLastPictureUpdate", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) Save(user *model.User) (*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.Save(user)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) Search(teamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.Search(teamId, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.Search", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.SearchInChannel(channelId, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.SearchInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) SearchInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.SearchInGroup(groupID, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.SearchInGroup", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.SearchNotInChannel(teamId, channelId, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.SearchNotInChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.SearchNotInTeam(notInTeamId, term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.SearchNotInTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.SearchWithoutTeam(term, options)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.SearchWithoutTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) Update(user *model.User, allowRoleUpdate bool) (*model.UserUpdate, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.Update(user, allowRoleUpdate)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.Update", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) (string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.UpdateAuthData(userId, service, authData, email, resetMfa)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateAuthData", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.UpdateFailedPasswordAttempts(userId, attempts)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateFailedPasswordAttempts", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) UpdateLastPictureUpdate(userId string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.UpdateLastPictureUpdate(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateLastPictureUpdate", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) UpdateMfaActive(userId string, active bool) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.UpdateMfaActive(userId, active)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateMfaActive", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) UpdateMfaSecret(userId string, secret string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.UpdateMfaSecret(userId, secret)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateMfaSecret", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) UpdatePassword(userId string, newPassword string) *model.AppError {
	start := timemodule.Now()

	err := s.UserStore.UpdatePassword(userId, newPassword)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdatePassword", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserStore) UpdateUpdateAt(userId string) (int64, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.UpdateUpdateAt(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.UpdateUpdateAt", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserStore) VerifyEmail(userId string, email string) (string, *model.AppError) {
	start := timemodule.Now()

	result, err := s.UserStore.VerifyEmail(userId, email)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserStore.VerifyEmail", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) Delete(tokenId string) error {
	start := timemodule.Now()

	err := s.UserAccessTokenStore.Delete(tokenId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserAccessTokenStore) DeleteAllForUser(userId string) error {
	start := timemodule.Now()

	err := s.UserAccessTokenStore.DeleteAllForUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.DeleteAllForUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserAccessTokenStore) Get(tokenId string) (*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.Get(tokenId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.Get", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) GetAll(offset int, limit int) ([]*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.GetAll(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.GetAll", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) GetByToken(tokenString string) (*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.GetByToken(tokenString)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.GetByToken", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) GetByUser(userId string, page int, perPage int) ([]*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.GetByUser(userId, page, perPage)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.GetByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) Save(token *model.UserAccessToken) (*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.Save(token)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) Search(term string) ([]*model.UserAccessToken, error) {
	start := timemodule.Now()

	result, err := s.UserAccessTokenStore.Search(term)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.Search", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserAccessTokenStore) UpdateTokenDisable(tokenId string) error {
	start := timemodule.Now()

	err := s.UserAccessTokenStore.UpdateTokenDisable(tokenId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.UpdateTokenDisable", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserAccessTokenStore) UpdateTokenEnable(tokenId string) error {
	start := timemodule.Now()

	err := s.UserAccessTokenStore.UpdateTokenEnable(tokenId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserAccessTokenStore.UpdateTokenEnable", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) error {
	start := timemodule.Now()

	err := s.UserTermsOfServiceStore.Delete(userId, termsOfServiceId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserTermsOfServiceStore.Delete", success, elapsed)
	}
	return err
}

func (s *TimerLayerUserTermsOfServiceStore) GetByUser(userId string) (*model.UserTermsOfService, error) {
	start := timemodule.Now()

	result, err := s.UserTermsOfServiceStore.GetByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserTermsOfServiceStore.GetByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerUserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) (*model.UserTermsOfService, error) {
	start := timemodule.Now()

	result, err := s.UserTermsOfServiceStore.Save(userTermsOfService)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("UserTermsOfServiceStore.Save", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) AnalyticsIncomingCount(teamId string) (int64, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.AnalyticsIncomingCount(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.AnalyticsIncomingCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.AnalyticsOutgoingCount(teamId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.AnalyticsOutgoingCount", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) ClearCaches() {
	start := timemodule.Now()

	s.WebhookStore.ClearCaches()

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.ClearCaches", success, elapsed)
	}
}

func (s *TimerLayerWebhookStore) DeleteIncoming(webhookId string, time int64) error {
	start := timemodule.Now()

	err := s.WebhookStore.DeleteIncoming(webhookId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.DeleteIncoming", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) DeleteOutgoing(webhookId string, time int64) error {
	start := timemodule.Now()

	err := s.WebhookStore.DeleteOutgoing(webhookId, time)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.DeleteOutgoing", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncoming(id, allowFromCache)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncoming", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncomingByChannel(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncomingByChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncomingByTeam(teamId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncomingByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetIncomingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncomingByTeamByUser(teamId, userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncomingByTeamByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncomingList(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncomingList", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetIncomingListByUser(userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetIncomingListByUser(userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetIncomingListByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoing(id)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingByChannel(channelId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingByChannel", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingByChannelByUser(channelId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingByChannelByUser(channelId, userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingByChannelByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingByTeam(teamId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingByTeam", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingByTeamByUser(teamId, userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingByTeamByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingList(offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingList", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) GetOutgoingListByUser(userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.GetOutgoingListByUser(userId, offset, limit)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.GetOutgoingListByUser", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) InvalidateWebhookCache(webhook string) {
	start := timemodule.Now()

	s.WebhookStore.InvalidateWebhookCache(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if true {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.InvalidateWebhookCache", success, elapsed)
	}
}

func (s *TimerLayerWebhookStore) PermanentDeleteIncomingByChannel(channelId string) error {
	start := timemodule.Now()

	err := s.WebhookStore.PermanentDeleteIncomingByChannel(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.PermanentDeleteIncomingByChannel", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) PermanentDeleteIncomingByUser(userId string) error {
	start := timemodule.Now()

	err := s.WebhookStore.PermanentDeleteIncomingByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.PermanentDeleteIncomingByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) error {
	start := timemodule.Now()

	err := s.WebhookStore.PermanentDeleteOutgoingByChannel(channelId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.PermanentDeleteOutgoingByChannel", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) PermanentDeleteOutgoingByUser(userId string) error {
	start := timemodule.Now()

	err := s.WebhookStore.PermanentDeleteOutgoingByUser(userId)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.PermanentDeleteOutgoingByUser", success, elapsed)
	}
	return err
}

func (s *TimerLayerWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.SaveIncoming(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.SaveIncoming", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.SaveOutgoing(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.SaveOutgoing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.UpdateIncoming(webhook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.UpdateIncoming", success, elapsed)
	}
	return result, err
}

func (s *TimerLayerWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	start := timemodule.Now()

	result, err := s.WebhookStore.UpdateOutgoing(hook)

	elapsed := float64(timemodule.Since(start)) / float64(timemodule.Second)
	if s.Root.Metrics != nil {
		success := "false"
		if err == nil {
			success = "true"
		}
		s.Root.Metrics.ObserveStoreMethodDuration("WebhookStore.UpdateOutgoing", success, elapsed)
	}
	return result, err
}

func (s *TimerLayer) Close() {
	s.Store.Close()
}

func (s *TimerLayer) DropAllTables() {
	s.Store.DropAllTables()
}

func (s *TimerLayer) GetCurrentSchemaVersion() string {
	return s.Store.GetCurrentSchemaVersion()
}

func (s *TimerLayer) LockToMaster() {
	s.Store.LockToMaster()
}

func (s *TimerLayer) MarkSystemRanUnitTests() {
	s.Store.MarkSystemRanUnitTests()
}

func (s *TimerLayer) SetContext(context context.Context) {
	s.Store.SetContext(context)
}

func (s *TimerLayer) TotalMasterDbConnections() int {
	return s.Store.TotalMasterDbConnections()
}

func (s *TimerLayer) TotalReadDbConnections() int {
	return s.Store.TotalReadDbConnections()
}

func (s *TimerLayer) TotalSearchDbConnections() int {
	return s.Store.TotalSearchDbConnections()
}

func (s *TimerLayer) UnlockFromMaster() {
	s.Store.UnlockFromMaster()
}

func New(childStore store.Store, metrics einterfaces.MetricsInterface) *TimerLayer {
	newStore := TimerLayer{
		Store:   childStore,
		Metrics: metrics,
	}

	newStore.AuditStore = &TimerLayerAuditStore{AuditStore: childStore.Audit(), Root: &newStore}
	newStore.BotStore = &TimerLayerBotStore{BotStore: childStore.Bot(), Root: &newStore}
	newStore.ChannelStore = &TimerLayerChannelStore{ChannelStore: childStore.Channel(), Root: &newStore}
	newStore.ChannelMemberHistoryStore = &TimerLayerChannelMemberHistoryStore{ChannelMemberHistoryStore: childStore.ChannelMemberHistory(), Root: &newStore}
	newStore.ClusterDiscoveryStore = &TimerLayerClusterDiscoveryStore{ClusterDiscoveryStore: childStore.ClusterDiscovery(), Root: &newStore}
	newStore.CommandStore = &TimerLayerCommandStore{CommandStore: childStore.Command(), Root: &newStore}
	newStore.CommandWebhookStore = &TimerLayerCommandWebhookStore{CommandWebhookStore: childStore.CommandWebhook(), Root: &newStore}
	newStore.ComplianceStore = &TimerLayerComplianceStore{ComplianceStore: childStore.Compliance(), Root: &newStore}
	newStore.EmojiStore = &TimerLayerEmojiStore{EmojiStore: childStore.Emoji(), Root: &newStore}
	newStore.FileInfoStore = &TimerLayerFileInfoStore{FileInfoStore: childStore.FileInfo(), Root: &newStore}
	newStore.GroupStore = &TimerLayerGroupStore{GroupStore: childStore.Group(), Root: &newStore}
	newStore.JobStore = &TimerLayerJobStore{JobStore: childStore.Job(), Root: &newStore}
	newStore.LicenseStore = &TimerLayerLicenseStore{LicenseStore: childStore.License(), Root: &newStore}
	newStore.LinkMetadataStore = &TimerLayerLinkMetadataStore{LinkMetadataStore: childStore.LinkMetadata(), Root: &newStore}
	newStore.OAuthStore = &TimerLayerOAuthStore{OAuthStore: childStore.OAuth(), Root: &newStore}
	newStore.PluginStore = &TimerLayerPluginStore{PluginStore: childStore.Plugin(), Root: &newStore}
	newStore.PostStore = &TimerLayerPostStore{PostStore: childStore.Post(), Root: &newStore}
	newStore.PreferenceStore = &TimerLayerPreferenceStore{PreferenceStore: childStore.Preference(), Root: &newStore}
	newStore.ReactionStore = &TimerLayerReactionStore{ReactionStore: childStore.Reaction(), Root: &newStore}
	newStore.RoleStore = &TimerLayerRoleStore{RoleStore: childStore.Role(), Root: &newStore}
	newStore.SchemeStore = &TimerLayerSchemeStore{SchemeStore: childStore.Scheme(), Root: &newStore}
	newStore.SessionStore = &TimerLayerSessionStore{SessionStore: childStore.Session(), Root: &newStore}
	newStore.StatusStore = &TimerLayerStatusStore{StatusStore: childStore.Status(), Root: &newStore}
	newStore.SystemStore = &TimerLayerSystemStore{SystemStore: childStore.System(), Root: &newStore}
	newStore.TeamStore = &TimerLayerTeamStore{TeamStore: childStore.Team(), Root: &newStore}
	newStore.TermsOfServiceStore = &TimerLayerTermsOfServiceStore{TermsOfServiceStore: childStore.TermsOfService(), Root: &newStore}
	newStore.TokenStore = &TimerLayerTokenStore{TokenStore: childStore.Token(), Root: &newStore}
	newStore.UserStore = &TimerLayerUserStore{UserStore: childStore.User(), Root: &newStore}
	newStore.UserAccessTokenStore = &TimerLayerUserAccessTokenStore{UserAccessTokenStore: childStore.UserAccessToken(), Root: &newStore}
	newStore.UserTermsOfServiceStore = &TimerLayerUserTermsOfServiceStore{UserTermsOfServiceStore: childStore.UserTermsOfService(), Root: &newStore}
	newStore.WebhookStore = &TimerLayerWebhookStore{WebhookStore: childStore.Webhook(), Root: &newStore}
	return &newStore
}
