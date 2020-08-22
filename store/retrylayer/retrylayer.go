// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package retrylayer

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/pkg/errors"
)

const mySQLDeadlockCode = uint16(1213)

type RetryLayer struct {
	store.Store
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

func (s *RetryLayer) Audit() store.AuditStore {
	return s.AuditStore
}

func (s *RetryLayer) Bot() store.BotStore {
	return s.BotStore
}

func (s *RetryLayer) Channel() store.ChannelStore {
	return s.ChannelStore
}

func (s *RetryLayer) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.ChannelMemberHistoryStore
}

func (s *RetryLayer) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.ClusterDiscoveryStore
}

func (s *RetryLayer) Command() store.CommandStore {
	return s.CommandStore
}

func (s *RetryLayer) CommandWebhook() store.CommandWebhookStore {
	return s.CommandWebhookStore
}

func (s *RetryLayer) Compliance() store.ComplianceStore {
	return s.ComplianceStore
}

func (s *RetryLayer) Emoji() store.EmojiStore {
	return s.EmojiStore
}

func (s *RetryLayer) FileInfo() store.FileInfoStore {
	return s.FileInfoStore
}

func (s *RetryLayer) Group() store.GroupStore {
	return s.GroupStore
}

func (s *RetryLayer) Job() store.JobStore {
	return s.JobStore
}

func (s *RetryLayer) License() store.LicenseStore {
	return s.LicenseStore
}

func (s *RetryLayer) LinkMetadata() store.LinkMetadataStore {
	return s.LinkMetadataStore
}

func (s *RetryLayer) OAuth() store.OAuthStore {
	return s.OAuthStore
}

func (s *RetryLayer) Plugin() store.PluginStore {
	return s.PluginStore
}

func (s *RetryLayer) Post() store.PostStore {
	return s.PostStore
}

func (s *RetryLayer) Preference() store.PreferenceStore {
	return s.PreferenceStore
}

func (s *RetryLayer) Reaction() store.ReactionStore {
	return s.ReactionStore
}

func (s *RetryLayer) Role() store.RoleStore {
	return s.RoleStore
}

func (s *RetryLayer) Scheme() store.SchemeStore {
	return s.SchemeStore
}

func (s *RetryLayer) Session() store.SessionStore {
	return s.SessionStore
}

func (s *RetryLayer) Status() store.StatusStore {
	return s.StatusStore
}

func (s *RetryLayer) System() store.SystemStore {
	return s.SystemStore
}

func (s *RetryLayer) Team() store.TeamStore {
	return s.TeamStore
}

func (s *RetryLayer) TermsOfService() store.TermsOfServiceStore {
	return s.TermsOfServiceStore
}

func (s *RetryLayer) Token() store.TokenStore {
	return s.TokenStore
}

func (s *RetryLayer) User() store.UserStore {
	return s.UserStore
}

func (s *RetryLayer) UserAccessToken() store.UserAccessTokenStore {
	return s.UserAccessTokenStore
}

func (s *RetryLayer) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.UserTermsOfServiceStore
}

func (s *RetryLayer) Webhook() store.WebhookStore {
	return s.WebhookStore
}

type RetryLayerAuditStore struct {
	store.AuditStore
	Root *RetryLayer
}

type RetryLayerBotStore struct {
	store.BotStore
	Root *RetryLayer
}

type RetryLayerChannelStore struct {
	store.ChannelStore
	Root *RetryLayer
}

type RetryLayerChannelMemberHistoryStore struct {
	store.ChannelMemberHistoryStore
	Root *RetryLayer
}

type RetryLayerClusterDiscoveryStore struct {
	store.ClusterDiscoveryStore
	Root *RetryLayer
}

type RetryLayerCommandStore struct {
	store.CommandStore
	Root *RetryLayer
}

type RetryLayerCommandWebhookStore struct {
	store.CommandWebhookStore
	Root *RetryLayer
}

type RetryLayerComplianceStore struct {
	store.ComplianceStore
	Root *RetryLayer
}

type RetryLayerEmojiStore struct {
	store.EmojiStore
	Root *RetryLayer
}

type RetryLayerFileInfoStore struct {
	store.FileInfoStore
	Root *RetryLayer
}

type RetryLayerGroupStore struct {
	store.GroupStore
	Root *RetryLayer
}

type RetryLayerJobStore struct {
	store.JobStore
	Root *RetryLayer
}

type RetryLayerLicenseStore struct {
	store.LicenseStore
	Root *RetryLayer
}

type RetryLayerLinkMetadataStore struct {
	store.LinkMetadataStore
	Root *RetryLayer
}

type RetryLayerOAuthStore struct {
	store.OAuthStore
	Root *RetryLayer
}

type RetryLayerPluginStore struct {
	store.PluginStore
	Root *RetryLayer
}

type RetryLayerPostStore struct {
	store.PostStore
	Root *RetryLayer
}

type RetryLayerPreferenceStore struct {
	store.PreferenceStore
	Root *RetryLayer
}

type RetryLayerReactionStore struct {
	store.ReactionStore
	Root *RetryLayer
}

type RetryLayerRoleStore struct {
	store.RoleStore
	Root *RetryLayer
}

type RetryLayerSchemeStore struct {
	store.SchemeStore
	Root *RetryLayer
}

type RetryLayerSessionStore struct {
	store.SessionStore
	Root *RetryLayer
}

type RetryLayerStatusStore struct {
	store.StatusStore
	Root *RetryLayer
}

type RetryLayerSystemStore struct {
	store.SystemStore
	Root *RetryLayer
}

type RetryLayerTeamStore struct {
	store.TeamStore
	Root *RetryLayer
}

type RetryLayerTermsOfServiceStore struct {
	store.TermsOfServiceStore
	Root *RetryLayer
}

type RetryLayerTokenStore struct {
	store.TokenStore
	Root *RetryLayer
}

type RetryLayerUserStore struct {
	store.UserStore
	Root *RetryLayer
}

type RetryLayerUserAccessTokenStore struct {
	store.UserAccessTokenStore
	Root *RetryLayer
}

type RetryLayerUserTermsOfServiceStore struct {
	store.UserTermsOfServiceStore
	Root *RetryLayer
}

type RetryLayerWebhookStore struct {
	store.WebhookStore
	Root *RetryLayer
}

func isRepeatableError(err error) bool {
	var pqErr *pq.Error
	var mysqlErr *mysql.MySQLError
	switch {
	case errors.As(errors.Cause(err), &pqErr):
		if pqErr.Code == "40001" || pqErr.Code == "40P01" {
			return true
		}
	case errors.As(errors.Cause(err), &mysqlErr):
		if mysqlErr.Number == mySQLDeadlockCode {
			return true
		}
	}
	return false
}

func (s *RetryLayerAuditStore) Get(user_id string, offset int, limit int) (model.Audits, error) {

	tries := 0
	for {
		result, err := s.AuditStore.Get(user_id, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerAuditStore) PermanentDeleteByUser(userId string) error {

	tries := 0
	for {
		err := s.AuditStore.PermanentDeleteByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerAuditStore) Save(audit *model.Audit) error {

	tries := 0
	for {
		err := s.AuditStore.Save(audit)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerBotStore) Get(userId string, includeDeleted bool) (*model.Bot, error) {

	tries := 0
	for {
		result, err := s.BotStore.Get(userId, includeDeleted)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerBotStore) GetAll(options *model.BotGetOptions) ([]*model.Bot, error) {

	tries := 0
	for {
		result, err := s.BotStore.GetAll(options)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerBotStore) PermanentDelete(userId string) error {

	tries := 0
	for {
		err := s.BotStore.PermanentDelete(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerBotStore) Save(bot *model.Bot) (*model.Bot, error) {

	tries := 0
	for {
		result, err := s.BotStore.Save(bot)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerBotStore) Update(bot *model.Bot) (*model.Bot, error) {

	tries := 0
	for {
		result, err := s.BotStore.Update(bot)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) (int64, *model.AppError) {

	return s.ChannelStore.AnalyticsDeletedTypeCount(teamId, channelType)

}

func (s *RetryLayerChannelStore) AnalyticsTypeCount(teamId string, channelType string) (int64, *model.AppError) {

	return s.ChannelStore.AnalyticsTypeCount(teamId, channelType)

}

func (s *RetryLayerChannelStore) AutocompleteInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.AutocompleteInTeam(teamId, term, includeDeleted)

}

func (s *RetryLayerChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.AutocompleteInTeamForSearch(teamId, userId, term, includeDeleted)

}

func (s *RetryLayerChannelStore) ClearAllCustomRoleAssignments() *model.AppError {

	return s.ChannelStore.ClearAllCustomRoleAssignments()

}

func (s *RetryLayerChannelStore) ClearCaches() {

	s.ChannelStore.ClearCaches()

}

func (s *RetryLayerChannelStore) ClearSidebarOnTeamLeave(userId string, teamId string) *model.AppError {

	return s.ChannelStore.ClearSidebarOnTeamLeave(userId, teamId)

}

func (s *RetryLayerChannelStore) CountPostsAfter(channelId string, timestamp int64, userId string) (int, *model.AppError) {

	return s.ChannelStore.CountPostsAfter(channelId, timestamp, userId)

}

func (s *RetryLayerChannelStore) CreateDirectChannel(userId *model.User, otherUserId *model.User) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.CreateDirectChannel(userId, otherUserId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) CreateInitialSidebarCategories(userId string, teamId string) error {

	tries := 0
	for {
		err := s.ChannelStore.CreateInitialSidebarCategories(userId, teamId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) CreateSidebarCategory(userId string, teamId string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.AppError) {

	return s.ChannelStore.CreateSidebarCategory(userId, teamId, newCategory)

}

func (s *RetryLayerChannelStore) Delete(channelId string, time int64) error {

	tries := 0
	for {
		err := s.ChannelStore.Delete(channelId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) DeleteSidebarCategory(categoryId string) *model.AppError {

	return s.ChannelStore.DeleteSidebarCategory(categoryId)

}

func (s *RetryLayerChannelStore) DeleteSidebarChannelsByPreferences(preferences *model.Preferences) error {

	tries := 0
	for {
		err := s.ChannelStore.DeleteSidebarChannelsByPreferences(preferences)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) Get(id string, allowFromCache bool) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.Get(id, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetAll(teamId string) ([]*model.Channel, *model.AppError) {

	return s.ChannelStore.GetAll(teamId)

}

func (s *RetryLayerChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, *model.AppError) {

	return s.ChannelStore.GetAllChannelMembersForUser(userId, allowFromCache, includeDeleted)

}

func (s *RetryLayerChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, *model.AppError) {

	return s.ChannelStore.GetAllChannelMembersNotifyPropsForChannel(channelId, allowFromCache)

}

func (s *RetryLayerChannelStore) GetAllChannels(page int, perPage int, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetAllChannels(page, perPage, opts)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetAllChannelsCount(opts store.ChannelSearchOpts) (int64, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetAllChannelsCount(opts)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) ([]*model.ChannelForExport, *model.AppError) {

	return s.ChannelStore.GetAllChannelsForExportAfter(limit, afterId)

}

func (s *RetryLayerChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) ([]*model.DirectChannelForExport, *model.AppError) {

	return s.ChannelStore.GetAllDirectChannelsForExportAfter(limit, afterId)

}

func (s *RetryLayerChannelStore) GetByName(team_id string, name string, allowFromCache bool) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetByName(team_id, name, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetByNameIncludeDeleted(team_id, name, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) ([]*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetByNames(team_id, names, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetChannelCounts(teamId string, userId string) (*model.ChannelCounts, *model.AppError) {

	return s.ChannelStore.GetChannelCounts(teamId, userId)

}

func (s *RetryLayerChannelStore) GetChannelMembersForExport(userId string, teamId string) ([]*model.ChannelMemberForExport, *model.AppError) {

	return s.ChannelStore.GetChannelMembersForExport(userId, teamId)

}

func (s *RetryLayerChannelStore) GetChannelMembersTimezones(channelId string) ([]model.StringMap, *model.AppError) {

	return s.ChannelStore.GetChannelMembersTimezones(channelId)

}

func (s *RetryLayerChannelStore) GetChannelUnread(channelId string, userId string) (*model.ChannelUnread, *model.AppError) {

	return s.ChannelStore.GetChannelUnread(channelId, userId)

}

func (s *RetryLayerChannelStore) GetChannels(teamId string, userId string, includeDeleted bool, lastDeleteAt int) (*model.ChannelList, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetChannels(teamId, userId, includeDeleted, lastDeleteAt)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetChannelsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.Channel, *model.AppError) {

	return s.ChannelStore.GetChannelsBatchForIndexing(startTime, endTime, limit)

}

func (s *RetryLayerChannelStore) GetChannelsByIds(channelIds []string, includeDeleted bool) ([]*model.Channel, *model.AppError) {

	return s.ChannelStore.GetChannelsByIds(channelIds, includeDeleted)

}

func (s *RetryLayerChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) (model.ChannelList, *model.AppError) {

	return s.ChannelStore.GetChannelsByScheme(schemeId, offset, limit)

}

func (s *RetryLayerChannelStore) GetDeleted(team_id string, offset int, limit int, userId string) (*model.ChannelList, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetDeleted(team_id, offset, limit, userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetDeletedByName(team_id string, name string) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetDeletedByName(team_id, name)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetForPost(postId string) (*model.Channel, *model.AppError) {

	return s.ChannelStore.GetForPost(postId)

}

func (s *RetryLayerChannelStore) GetFromMaster(id string) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetFromMaster(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetGuestCount(channelId string, allowFromCache bool) (int64, *model.AppError) {

	return s.ChannelStore.GetGuestCount(channelId, allowFromCache)

}

func (s *RetryLayerChannelStore) GetMember(channelId string, userId string) (*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.GetMember(channelId, userId)

}

func (s *RetryLayerChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, *model.AppError) {

	return s.ChannelStore.GetMemberCount(channelId, allowFromCache)

}

func (s *RetryLayerChannelStore) GetMemberCountFromCache(channelId string) int64 {

	return s.ChannelStore.GetMemberCountFromCache(channelId)

}

func (s *RetryLayerChannelStore) GetMemberCountsByGroup(channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, *model.AppError) {

	return s.ChannelStore.GetMemberCountsByGroup(channelID, includeTimezones)

}

func (s *RetryLayerChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.GetMemberForPost(postId, userId)

}

func (s *RetryLayerChannelStore) GetMembers(channelId string, offset int, limit int) (*model.ChannelMembers, *model.AppError) {

	return s.ChannelStore.GetMembers(channelId, offset, limit)

}

func (s *RetryLayerChannelStore) GetMembersByIds(channelId string, userIds []string) (*model.ChannelMembers, *model.AppError) {

	return s.ChannelStore.GetMembersByIds(channelId, userIds)

}

func (s *RetryLayerChannelStore) GetMembersForUser(teamId string, userId string) (*model.ChannelMembers, *model.AppError) {

	return s.ChannelStore.GetMembersForUser(teamId, userId)

}

func (s *RetryLayerChannelStore) GetMembersForUserWithPagination(teamId string, userId string, page int, perPage int) (*model.ChannelMembers, *model.AppError) {

	return s.ChannelStore.GetMembersForUserWithPagination(teamId, userId, page, perPage)

}

func (s *RetryLayerChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) (*model.ChannelList, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.GetMoreChannels(teamId, userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, *model.AppError) {

	return s.ChannelStore.GetPinnedPostCount(channelId, allowFromCache)

}

func (s *RetryLayerChannelStore) GetPinnedPosts(channelId string) (*model.PostList, *model.AppError) {

	return s.ChannelStore.GetPinnedPosts(channelId)

}

func (s *RetryLayerChannelStore) GetPrivateChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.GetPrivateChannelsForTeam(teamId, offset, limit)

}

func (s *RetryLayerChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.GetPublicChannelsByIdsForTeam(teamId, channelIds)

}

func (s *RetryLayerChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.GetPublicChannelsForTeam(teamId, offset, limit)

}

func (s *RetryLayerChannelStore) GetSidebarCategories(userId string, teamId string) (*model.OrderedSidebarCategories, *model.AppError) {

	return s.ChannelStore.GetSidebarCategories(userId, teamId)

}

func (s *RetryLayerChannelStore) GetSidebarCategory(categoryId string) (*model.SidebarCategoryWithChannels, *model.AppError) {

	return s.ChannelStore.GetSidebarCategory(categoryId)

}

func (s *RetryLayerChannelStore) GetSidebarCategoryOrder(userId string, teamId string) ([]string, *model.AppError) {

	return s.ChannelStore.GetSidebarCategoryOrder(userId, teamId)

}

func (s *RetryLayerChannelStore) GetTeamChannels(teamId string) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.GetTeamChannels(teamId)

}

func (s *RetryLayerChannelStore) GroupSyncedChannelCount() (int64, *model.AppError) {

	return s.ChannelStore.GroupSyncedChannelCount()

}

func (s *RetryLayerChannelStore) IncrementMentionCount(channelId string, userId string) *model.AppError {

	return s.ChannelStore.IncrementMentionCount(channelId, userId)

}

func (s *RetryLayerChannelStore) InvalidateAllChannelMembersForUser(userId string) {

	s.ChannelStore.InvalidateAllChannelMembersForUser(userId)

}

func (s *RetryLayerChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {

	s.ChannelStore.InvalidateCacheForChannelMembersNotifyProps(channelId)

}

func (s *RetryLayerChannelStore) InvalidateChannel(id string) {

	s.ChannelStore.InvalidateChannel(id)

}

func (s *RetryLayerChannelStore) InvalidateChannelByName(teamId string, name string) {

	s.ChannelStore.InvalidateChannelByName(teamId, name)

}

func (s *RetryLayerChannelStore) InvalidateGuestCount(channelId string) {

	s.ChannelStore.InvalidateGuestCount(channelId)

}

func (s *RetryLayerChannelStore) InvalidateMemberCount(channelId string) {

	s.ChannelStore.InvalidateMemberCount(channelId)

}

func (s *RetryLayerChannelStore) InvalidatePinnedPostCount(channelId string) {

	s.ChannelStore.InvalidatePinnedPostCount(channelId)

}

func (s *RetryLayerChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {

	return s.ChannelStore.IsUserInChannelUseCache(userId, channelId)

}

func (s *RetryLayerChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) (map[string]string, *model.AppError) {

	return s.ChannelStore.MigrateChannelMembers(fromChannelId, fromUserId)

}

func (s *RetryLayerChannelStore) MigratePublicChannels() error {

	tries := 0
	for {
		err := s.ChannelStore.MigratePublicChannels()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) PermanentDelete(channelId string) error {

	tries := 0
	for {
		err := s.ChannelStore.PermanentDelete(channelId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) PermanentDeleteByTeam(teamId string) error {

	tries := 0
	for {
		err := s.ChannelStore.PermanentDeleteByTeam(teamId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) PermanentDeleteMembersByChannel(channelId string) *model.AppError {

	return s.ChannelStore.PermanentDeleteMembersByChannel(channelId)

}

func (s *RetryLayerChannelStore) PermanentDeleteMembersByUser(userId string) *model.AppError {

	return s.ChannelStore.PermanentDeleteMembersByUser(userId)

}

func (s *RetryLayerChannelStore) RemoveAllDeactivatedMembers(channelId string) *model.AppError {

	return s.ChannelStore.RemoveAllDeactivatedMembers(channelId)

}

func (s *RetryLayerChannelStore) RemoveMember(channelId string, userId string) *model.AppError {

	return s.ChannelStore.RemoveMember(channelId, userId)

}

func (s *RetryLayerChannelStore) RemoveMembers(channelId string, userIds []string) *model.AppError {

	return s.ChannelStore.RemoveMembers(channelId, userIds)

}

func (s *RetryLayerChannelStore) ResetAllChannelSchemes() *model.AppError {

	return s.ChannelStore.ResetAllChannelSchemes()

}

func (s *RetryLayerChannelStore) Restore(channelId string, time int64) error {

	tries := 0
	for {
		err := s.ChannelStore.Restore(channelId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.Save(channel, maxChannelsPerTeam)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.SaveDirectChannel(channel, member1, member2)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) SaveMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.SaveMember(member)

}

func (s *RetryLayerChannelStore) SaveMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.SaveMultipleMembers(members)

}

func (s *RetryLayerChannelStore) SearchAllChannels(term string, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, int64, *model.AppError) {

	return s.ChannelStore.SearchAllChannels(term, opts)

}

func (s *RetryLayerChannelStore) SearchArchivedInTeam(teamId string, term string, userId string) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.SearchArchivedInTeam(teamId, term, userId)

}

func (s *RetryLayerChannelStore) SearchForUserInTeam(userId string, teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.SearchForUserInTeam(userId, teamId, term, includeDeleted)

}

func (s *RetryLayerChannelStore) SearchGroupChannels(userId string, term string) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.SearchGroupChannels(userId, term)

}

func (s *RetryLayerChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.SearchInTeam(teamId, term, includeDeleted)

}

func (s *RetryLayerChannelStore) SearchMore(userId string, teamId string, term string) (*model.ChannelList, *model.AppError) {

	return s.ChannelStore.SearchMore(userId, teamId, term)

}

func (s *RetryLayerChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) error {

	tries := 0
	for {
		err := s.ChannelStore.SetDeleteAt(channelId, deleteAt, updateAt)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) Update(channel *model.Channel) (*model.Channel, error) {

	tries := 0
	for {
		result, err := s.ChannelStore.Update(channel)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelStore) UpdateLastViewedAt(channelIds []string, userId string) (map[string]int64, *model.AppError) {

	return s.ChannelStore.UpdateLastViewedAt(channelIds, userId)

}

func (s *RetryLayerChannelStore) UpdateLastViewedAtPost(unreadPost *model.Post, userID string, mentionCount int) (*model.ChannelUnreadAt, *model.AppError) {

	return s.ChannelStore.UpdateLastViewedAtPost(unreadPost, userID, mentionCount)

}

func (s *RetryLayerChannelStore) UpdateMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.UpdateMember(member)

}

func (s *RetryLayerChannelStore) UpdateMembersRole(channelID string, userIDs []string) *model.AppError {

	return s.ChannelStore.UpdateMembersRole(channelID, userIDs)

}

func (s *RetryLayerChannelStore) UpdateMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {

	return s.ChannelStore.UpdateMultipleMembers(members)

}

func (s *RetryLayerChannelStore) UpdateSidebarCategories(userId string, teamId string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, *model.AppError) {

	return s.ChannelStore.UpdateSidebarCategories(userId, teamId, categories)

}

func (s *RetryLayerChannelStore) UpdateSidebarCategoryOrder(userId string, teamId string, categoryOrder []string) *model.AppError {

	return s.ChannelStore.UpdateSidebarCategoryOrder(userId, teamId, categoryOrder)

}

func (s *RetryLayerChannelStore) UpdateSidebarChannelCategoryOnMove(channel *model.Channel, newTeamId string) *model.AppError {

	return s.ChannelStore.UpdateSidebarChannelCategoryOnMove(channel, newTeamId)

}

func (s *RetryLayerChannelStore) UpdateSidebarChannelsByPreferences(preferences *model.Preferences) error {

	tries := 0
	for {
		err := s.ChannelStore.UpdateSidebarChannelsByPreferences(preferences)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelStore) UserBelongsToChannels(userId string, channelIds []string) (bool, *model.AppError) {

	return s.ChannelStore.UserBelongsToChannels(userId, channelIds)

}

func (s *RetryLayerChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) ([]*model.ChannelMemberHistoryResult, error) {

	tries := 0
	for {
		result, err := s.ChannelMemberHistoryStore.GetUsersInChannelDuring(startTime, endTime, channelId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) error {

	tries := 0
	for {
		err := s.ChannelMemberHistoryStore.LogJoinEvent(userId, channelId, joinTime)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) error {

	tries := 0
	for {
		err := s.ChannelMemberHistoryStore.LogLeaveEvent(userId, channelId, leaveTime)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {

	tries := 0
	for {
		result, err := s.ChannelMemberHistoryStore.PermanentDeleteBatch(endTime, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) Cleanup() error {

	tries := 0
	for {
		err := s.ClusterDiscoveryStore.Cleanup()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, error) {

	tries := 0
	for {
		result, err := s.ClusterDiscoveryStore.Delete(discovery)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, error) {

	tries := 0
	for {
		result, err := s.ClusterDiscoveryStore.Exists(discovery)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, error) {

	tries := 0
	for {
		result, err := s.ClusterDiscoveryStore.GetAll(discoveryType, clusterName)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) error {

	tries := 0
	for {
		err := s.ClusterDiscoveryStore.Save(discovery)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) error {

	tries := 0
	for {
		err := s.ClusterDiscoveryStore.SetLastPingAt(discovery)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerCommandStore) AnalyticsCommandCount(teamId string) (int64, error) {

	tries := 0
	for {
		result, err := s.CommandStore.AnalyticsCommandCount(teamId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandStore) Delete(commandId string, time int64) error {

	tries := 0
	for {
		err := s.CommandStore.Delete(commandId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerCommandStore) Get(id string) (*model.Command, error) {

	tries := 0
	for {
		result, err := s.CommandStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandStore) GetByTeam(teamId string) ([]*model.Command, error) {

	tries := 0
	for {
		result, err := s.CommandStore.GetByTeam(teamId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandStore) GetByTrigger(teamId string, trigger string) (*model.Command, error) {

	tries := 0
	for {
		result, err := s.CommandStore.GetByTrigger(teamId, trigger)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandStore) PermanentDeleteByTeam(teamId string) error {

	tries := 0
	for {
		err := s.CommandStore.PermanentDeleteByTeam(teamId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerCommandStore) PermanentDeleteByUser(userId string) error {

	tries := 0
	for {
		err := s.CommandStore.PermanentDeleteByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerCommandStore) Save(webhook *model.Command) (*model.Command, error) {

	tries := 0
	for {
		result, err := s.CommandStore.Save(webhook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandStore) Update(hook *model.Command) (*model.Command, error) {

	tries := 0
	for {
		result, err := s.CommandStore.Update(hook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandWebhookStore) Cleanup() {

	s.CommandWebhookStore.Cleanup()

}

func (s *RetryLayerCommandWebhookStore) Get(id string) (*model.CommandWebhook, error) {

	tries := 0
	for {
		result, err := s.CommandWebhookStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandWebhookStore) Save(webhook *model.CommandWebhook) (*model.CommandWebhook, error) {

	tries := 0
	for {
		result, err := s.CommandWebhookStore.Save(webhook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerCommandWebhookStore) TryUse(id string, limit int) error {

	tries := 0
	for {
		err := s.CommandWebhookStore.TryUse(id, limit)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.ComplianceExport(compliance)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerComplianceStore) Get(id string) (*model.Compliance, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerComplianceStore) GetAll(offset int, limit int) (model.Compliances, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.GetAll(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.MessageExport(after, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.Save(compliance)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, error) {

	tries := 0
	for {
		result, err := s.ComplianceStore.Update(compliance)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) Delete(emoji *model.Emoji, time int64) error {

	tries := 0
	for {
		err := s.EmojiStore.Delete(emoji, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerEmojiStore) Get(id string, allowFromCache bool) (*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.Get(id, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) GetByName(name string, allowFromCache bool) (*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.GetByName(name, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) GetList(offset int, limit int, sort string) ([]*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.GetList(offset, limit, sort)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) GetMultipleByName(names []string) ([]*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.GetMultipleByName(names)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.Save(emoji)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerEmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {

	tries := 0
	for {
		result, err := s.EmojiStore.Search(name, prefixOnly, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) AttachToPost(fileId string, postId string, creatorId string) error {

	tries := 0
	for {
		err := s.FileInfoStore.AttachToPost(fileId, postId, creatorId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerFileInfoStore) ClearCaches() {

	s.FileInfoStore.ClearCaches()

}

func (s *RetryLayerFileInfoStore) DeleteForPost(postId string) (string, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.DeleteForPost(postId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) Get(id string) (*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) GetByPath(path string) (*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.GetByPath(path)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) GetForPost(postId string, readFromMaster bool, includeDeleted bool, allowFromCache bool) ([]*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.GetForPost(postId, readFromMaster, includeDeleted, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) GetForUser(userId string) ([]*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.GetForUser(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) GetWithOptions(page int, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.GetWithOptions(page, perPage, opt)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) InvalidateFileInfosForPostCache(postId string, deleted bool) {

	s.FileInfoStore.InvalidateFileInfosForPostCache(postId, deleted)

}

func (s *RetryLayerFileInfoStore) PermanentDelete(fileId string) error {

	tries := 0
	for {
		err := s.FileInfoStore.PermanentDelete(fileId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerFileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.PermanentDeleteBatch(endTime, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) PermanentDeleteByUser(userId string) (int64, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.PermanentDeleteByUser(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerFileInfoStore) Save(info *model.FileInfo) (*model.FileInfo, error) {

	tries := 0
	for {
		result, err := s.FileInfoStore.Save(info)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerGroupStore) AdminRoleGroupsForSyncableMember(userID string, syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {

	return s.GroupStore.AdminRoleGroupsForSyncableMember(userID, syncableID, syncableType)

}

func (s *RetryLayerGroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {

	return s.GroupStore.ChannelMembersMinusGroupMembers(channelID, groupIDs, page, perPage)

}

func (s *RetryLayerGroupStore) ChannelMembersToAdd(since int64, channelID *string) ([]*model.UserChannelIDPair, *model.AppError) {

	return s.GroupStore.ChannelMembersToAdd(since, channelID)

}

func (s *RetryLayerGroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, *model.AppError) {

	return s.GroupStore.ChannelMembersToRemove(channelID)

}

func (s *RetryLayerGroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, *model.AppError) {

	return s.GroupStore.CountChannelMembersMinusGroupMembers(channelID, groupIDs)

}

func (s *RetryLayerGroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, *model.AppError) {

	return s.GroupStore.CountGroupsByChannel(channelId, opts)

}

func (s *RetryLayerGroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, *model.AppError) {

	return s.GroupStore.CountGroupsByTeam(teamId, opts)

}

func (s *RetryLayerGroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, *model.AppError) {

	return s.GroupStore.CountTeamMembersMinusGroupMembers(teamID, groupIDs)

}

func (s *RetryLayerGroupStore) Create(group *model.Group) (*model.Group, *model.AppError) {

	return s.GroupStore.Create(group)

}

func (s *RetryLayerGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {

	return s.GroupStore.CreateGroupSyncable(groupSyncable)

}

func (s *RetryLayerGroupStore) Delete(groupID string) (*model.Group, *model.AppError) {

	return s.GroupStore.Delete(groupID)

}

func (s *RetryLayerGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {

	return s.GroupStore.DeleteGroupSyncable(groupID, syncableID, syncableType)

}

func (s *RetryLayerGroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {

	return s.GroupStore.DeleteMember(groupID, userID)

}

func (s *RetryLayerGroupStore) DistinctGroupMemberCount() (int64, *model.AppError) {

	return s.GroupStore.DistinctGroupMemberCount()

}

func (s *RetryLayerGroupStore) Get(groupID string) (*model.Group, *model.AppError) {

	return s.GroupStore.Get(groupID)

}

func (s *RetryLayerGroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {

	return s.GroupStore.GetAllBySource(groupSource)

}

func (s *RetryLayerGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, *model.AppError) {

	return s.GroupStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)

}

func (s *RetryLayerGroupStore) GetByIDs(groupIDs []string) ([]*model.Group, *model.AppError) {

	return s.GroupStore.GetByIDs(groupIDs)

}

func (s *RetryLayerGroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, *model.AppError) {

	return s.GroupStore.GetByName(name, opts)

}

func (s *RetryLayerGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, *model.AppError) {

	return s.GroupStore.GetByRemoteID(remoteID, groupSource)

}

func (s *RetryLayerGroupStore) GetByUser(userId string) ([]*model.Group, *model.AppError) {

	return s.GroupStore.GetByUser(userId)

}

func (s *RetryLayerGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {

	return s.GroupStore.GetGroupSyncable(groupID, syncableID, syncableType)

}

func (s *RetryLayerGroupStore) GetGroups(page int, perPage int, opts model.GroupSearchOpts) ([]*model.Group, *model.AppError) {

	return s.GroupStore.GetGroups(page, perPage, opts)

}

func (s *RetryLayerGroupStore) GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, *model.AppError) {

	return s.GroupStore.GetGroupsAssociatedToChannelsByTeam(teamId, opts)

}

func (s *RetryLayerGroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {

	return s.GroupStore.GetGroupsByChannel(channelId, opts)

}

func (s *RetryLayerGroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {

	return s.GroupStore.GetGroupsByTeam(teamId, opts)

}

func (s *RetryLayerGroupStore) GetMemberCount(groupID string) (int64, *model.AppError) {

	return s.GroupStore.GetMemberCount(groupID)

}

func (s *RetryLayerGroupStore) GetMemberUsers(groupID string) ([]*model.User, *model.AppError) {

	return s.GroupStore.GetMemberUsers(groupID)

}

func (s *RetryLayerGroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, *model.AppError) {

	return s.GroupStore.GetMemberUsersInTeam(groupID, teamID)

}

func (s *RetryLayerGroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, *model.AppError) {

	return s.GroupStore.GetMemberUsersNotInChannel(groupID, channelID)

}

func (s *RetryLayerGroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, *model.AppError) {

	return s.GroupStore.GetMemberUsersPage(groupID, page, perPage)

}

func (s *RetryLayerGroupStore) GroupChannelCount() (int64, *model.AppError) {

	return s.GroupStore.GroupChannelCount()

}

func (s *RetryLayerGroupStore) GroupCount() (int64, *model.AppError) {

	return s.GroupStore.GroupCount()

}

func (s *RetryLayerGroupStore) GroupCountWithAllowReference() (int64, *model.AppError) {

	return s.GroupStore.GroupCountWithAllowReference()

}

func (s *RetryLayerGroupStore) GroupMemberCount() (int64, *model.AppError) {

	return s.GroupStore.GroupMemberCount()

}

func (s *RetryLayerGroupStore) GroupTeamCount() (int64, *model.AppError) {

	return s.GroupStore.GroupTeamCount()

}

func (s *RetryLayerGroupStore) PermanentDeleteMembersByUser(userId string) *model.AppError {

	return s.GroupStore.PermanentDeleteMembersByUser(userId)

}

func (s *RetryLayerGroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {

	return s.GroupStore.PermittedSyncableAdmins(syncableID, syncableType)

}

func (s *RetryLayerGroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {

	return s.GroupStore.TeamMembersMinusGroupMembers(teamID, groupIDs, page, perPage)

}

func (s *RetryLayerGroupStore) TeamMembersToAdd(since int64, teamID *string) ([]*model.UserTeamIDPair, *model.AppError) {

	return s.GroupStore.TeamMembersToAdd(since, teamID)

}

func (s *RetryLayerGroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, *model.AppError) {

	return s.GroupStore.TeamMembersToRemove(teamID)

}

func (s *RetryLayerGroupStore) Update(group *model.Group) (*model.Group, *model.AppError) {

	return s.GroupStore.Update(group)

}

func (s *RetryLayerGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {

	return s.GroupStore.UpdateGroupSyncable(groupSyncable)

}

func (s *RetryLayerGroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {

	return s.GroupStore.UpsertMember(groupID, userID)

}

func (s *RetryLayerJobStore) Delete(id string) (string, error) {

	tries := 0
	for {
		result, err := s.JobStore.Delete(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) Get(id string) (*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetAllByStatus(status string) ([]*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetAllByStatus(status)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetAllByType(jobType string) ([]*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetAllByType(jobType)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetAllByTypePage(jobType string, offset int, limit int) ([]*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetAllByTypePage(jobType, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetAllPage(offset int, limit int) ([]*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetAllPage(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetCountByStatusAndType(status string, jobType string) (int64, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetCountByStatusAndType(status, jobType)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetNewestJobByStatusAndType(status string, jobType string) (*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetNewestJobByStatusAndType(status, jobType)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) GetNewestJobByStatusesAndType(statuses []string, jobType string) (*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.GetNewestJobByStatusesAndType(statuses, jobType)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) Save(job *model.Job) (*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.Save(job)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) UpdateOptimistically(job *model.Job, currentStatus string) (bool, error) {

	tries := 0
	for {
		result, err := s.JobStore.UpdateOptimistically(job, currentStatus)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) UpdateStatus(id string, status string) (*model.Job, error) {

	tries := 0
	for {
		result, err := s.JobStore.UpdateStatus(id, status)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerJobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) (bool, error) {

	tries := 0
	for {
		result, err := s.JobStore.UpdateStatusOptimistically(id, currentStatus, newStatus)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerLicenseStore) Get(id string) (*model.LicenseRecord, error) {

	tries := 0
	for {
		result, err := s.LicenseStore.Get(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, error) {

	tries := 0
	for {
		result, err := s.LicenseStore.Save(license)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerLinkMetadataStore) Get(url string, timestamp int64) (*model.LinkMetadata, error) {

	tries := 0
	for {
		result, err := s.LinkMetadataStore.Get(url, timestamp)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerLinkMetadataStore) Save(linkMetadata *model.LinkMetadata) (*model.LinkMetadata, error) {

	tries := 0
	for {
		result, err := s.LinkMetadataStore.Save(linkMetadata)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) DeleteApp(id string) error {

	tries := 0
	for {
		err := s.OAuthStore.DeleteApp(id)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAccessData(token string) (*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAccessData(token)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAccessDataByRefreshToken(token string) (*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAccessDataByRefreshToken(token)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAccessDataByUserForApp(userId string, clientId string) ([]*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAccessDataByUserForApp(userId, clientId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetApp(id string) (*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetApp(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAppByUser(userId string, offset int, limit int) ([]*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAppByUser(userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetApps(offset int, limit int) ([]*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetApps(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAuthData(code string) (*model.AuthData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAuthData(code)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetAuthorizedApps(userId string, offset int, limit int) ([]*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetAuthorizedApps(userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) GetPreviousAccessData(userId string, clientId string) (*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.GetPreviousAccessData(userId, clientId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) PermanentDeleteAuthDataByUser(userId string) error {

	tries := 0
	for {
		err := s.OAuthStore.PermanentDeleteAuthDataByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerOAuthStore) RemoveAccessData(token string) error {

	tries := 0
	for {
		err := s.OAuthStore.RemoveAccessData(token)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerOAuthStore) RemoveAllAccessData() error {

	tries := 0
	for {
		err := s.OAuthStore.RemoveAllAccessData()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerOAuthStore) RemoveAuthData(code string) error {

	tries := 0
	for {
		err := s.OAuthStore.RemoveAuthData(code)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerOAuthStore) SaveAccessData(accessData *model.AccessData) (*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.SaveAccessData(accessData)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) SaveApp(app *model.OAuthApp) (*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.SaveApp(app)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) SaveAuthData(authData *model.AuthData) (*model.AuthData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.SaveAuthData(authData)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) UpdateAccessData(accessData *model.AccessData) (*model.AccessData, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.UpdateAccessData(accessData)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerOAuthStore) UpdateApp(app *model.OAuthApp) (*model.OAuthApp, error) {

	tries := 0
	for {
		result, err := s.OAuthStore.UpdateApp(app)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) CompareAndDelete(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {

	tries := 0
	for {
		result, err := s.PluginStore.CompareAndDelete(keyVal, oldValue)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {

	tries := 0
	for {
		result, err := s.PluginStore.CompareAndSet(keyVal, oldValue)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) Delete(pluginId string, key string) error {

	tries := 0
	for {
		err := s.PluginStore.Delete(pluginId, key)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPluginStore) DeleteAllExpired() error {

	tries := 0
	for {
		err := s.PluginStore.DeleteAllExpired()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPluginStore) DeleteAllForPlugin(PluginId string) error {

	tries := 0
	for {
		err := s.PluginStore.DeleteAllForPlugin(PluginId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPluginStore) Get(pluginId string, key string) (*model.PluginKeyValue, error) {

	tries := 0
	for {
		result, err := s.PluginStore.Get(pluginId, key)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) List(pluginId string, page int, perPage int) ([]string, error) {

	tries := 0
	for {
		result, err := s.PluginStore.List(pluginId, page, perPage)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) (*model.PluginKeyValue, error) {

	tries := 0
	for {
		result, err := s.PluginStore.SaveOrUpdate(keyVal)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPluginStore) SetWithOptions(pluginId string, key string, value []byte, options model.PluginKVSetOptions) (bool, error) {

	tries := 0
	for {
		result, err := s.PluginStore.SetWithOptions(pluginId, key, value, options)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) (int64, *model.AppError) {

	return s.PostStore.AnalyticsPostCount(teamId, mustHaveFile, mustHaveHashtag)

}

func (s *RetryLayerPostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, *model.AppError) {

	return s.PostStore.AnalyticsPostCountsByDay(options)

}

func (s *RetryLayerPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) (model.AnalyticsRows, *model.AppError) {

	return s.PostStore.AnalyticsUserCountsWithPostsByDay(teamId)

}

func (s *RetryLayerPostStore) ClearCaches() {

	s.PostStore.ClearCaches()

}

func (s *RetryLayerPostStore) Delete(postId string, time int64, deleteByID string) error {

	tries := 0
	for {
		err := s.PostStore.Delete(postId, time, deleteByID)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPostStore) Get(id string, skipFetchThreads bool) (*model.PostList, error) {

	tries := 0
	for {
		result, err := s.PostStore.Get(id, skipFetchThreads)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) ([]*model.DirectPostForExport, *model.AppError) {

	return s.PostStore.GetDirectPostParentsForExportAfter(limit, afterId)

}

func (s *RetryLayerPostStore) GetEtag(channelId string, allowFromCache bool) string {

	return s.PostStore.GetEtag(channelId, allowFromCache)

}

func (s *RetryLayerPostStore) GetFlaggedPosts(userId string, offset int, limit int) (*model.PostList, *model.AppError) {

	return s.PostStore.GetFlaggedPosts(userId, offset, limit)

}

func (s *RetryLayerPostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) (*model.PostList, *model.AppError) {

	return s.PostStore.GetFlaggedPostsForChannel(userId, channelId, offset, limit)

}

func (s *RetryLayerPostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) (*model.PostList, *model.AppError) {

	return s.PostStore.GetFlaggedPostsForTeam(userId, teamId, offset, limit)

}

func (s *RetryLayerPostStore) GetMaxPostSize() int {

	return s.PostStore.GetMaxPostSize()

}

func (s *RetryLayerPostStore) GetOldest() (*model.Post, *model.AppError) {

	return s.PostStore.GetOldest()

}

func (s *RetryLayerPostStore) GetOldestEntityCreationTime() (int64, *model.AppError) {

	return s.PostStore.GetOldestEntityCreationTime()

}

func (s *RetryLayerPostStore) GetParentsForExportAfter(limit int, afterId string) ([]*model.PostForExport, *model.AppError) {

	return s.PostStore.GetParentsForExportAfter(limit, afterId)

}

func (s *RetryLayerPostStore) GetPostAfterTime(channelId string, time int64) (*model.Post, *model.AppError) {

	return s.PostStore.GetPostAfterTime(channelId, time)

}

func (s *RetryLayerPostStore) GetPostIdAfterTime(channelId string, time int64) (string, *model.AppError) {

	return s.PostStore.GetPostIdAfterTime(channelId, time)

}

func (s *RetryLayerPostStore) GetPostIdBeforeTime(channelId string, time int64) (string, *model.AppError) {

	return s.PostStore.GetPostIdBeforeTime(channelId, time)

}

func (s *RetryLayerPostStore) GetPosts(options model.GetPostsOptions, allowFromCache bool) (*model.PostList, *model.AppError) {

	return s.PostStore.GetPosts(options, allowFromCache)

}

func (s *RetryLayerPostStore) GetPostsAfter(options model.GetPostsOptions) (*model.PostList, *model.AppError) {

	return s.PostStore.GetPostsAfter(options)

}

func (s *RetryLayerPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.PostForIndexing, *model.AppError) {

	return s.PostStore.GetPostsBatchForIndexing(startTime, endTime, limit)

}

func (s *RetryLayerPostStore) GetPostsBefore(options model.GetPostsOptions) (*model.PostList, *model.AppError) {

	return s.PostStore.GetPostsBefore(options)

}

func (s *RetryLayerPostStore) GetPostsByIds(postIds []string) ([]*model.Post, *model.AppError) {

	return s.PostStore.GetPostsByIds(postIds)

}

func (s *RetryLayerPostStore) GetPostsCreatedAt(channelId string, time int64) ([]*model.Post, *model.AppError) {

	return s.PostStore.GetPostsCreatedAt(channelId, time)

}

func (s *RetryLayerPostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool) (*model.PostList, *model.AppError) {

	return s.PostStore.GetPostsSince(options, allowFromCache)

}

func (s *RetryLayerPostStore) GetRepliesForExport(parentId string) ([]*model.ReplyForExport, *model.AppError) {

	return s.PostStore.GetRepliesForExport(parentId)

}

func (s *RetryLayerPostStore) GetSingle(id string) (*model.Post, error) {

	tries := 0
	for {
		result, err := s.PostStore.GetSingle(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPostStore) InvalidateLastPostTimeCache(channelId string) {

	s.PostStore.InvalidateLastPostTimeCache(channelId)

}

func (s *RetryLayerPostStore) Overwrite(post *model.Post) (*model.Post, *model.AppError) {

	return s.PostStore.Overwrite(post)

}

func (s *RetryLayerPostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, *model.AppError) {

	return s.PostStore.OverwriteMultiple(posts)

}

func (s *RetryLayerPostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {

	return s.PostStore.PermanentDeleteBatch(endTime, limit)

}

func (s *RetryLayerPostStore) PermanentDeleteByChannel(channelId string) error {

	tries := 0
	for {
		err := s.PostStore.PermanentDeleteByChannel(channelId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPostStore) PermanentDeleteByUser(userId string) error {

	tries := 0
	for {
		err := s.PostStore.PermanentDeleteByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPostStore) Save(post *model.Post) (*model.Post, error) {

	tries := 0
	for {
		result, err := s.PostStore.Save(post)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {

	tries := 0
	for {
		result, resultVar1, err := s.PostStore.SaveMultiple(posts)
		if err == nil {
			return result, resultVar1, nil
		}
		if !isRepeatableError(err) {
			return result, resultVar1, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, resultVar1, err
		}
	}

}

func (s *RetryLayerPostStore) Search(teamId string, userId string, params *model.SearchParams) (*model.PostList, *model.AppError) {

	return s.PostStore.Search(teamId, userId, params)

}

func (s *RetryLayerPostStore) SearchPostsInTeamForUser(paramsList []*model.SearchParams, userId string, teamId string, isOrSearch bool, includeDeletedChannels bool, page int, perPage int) (*model.PostSearchResults, *model.AppError) {

	return s.PostStore.SearchPostsInTeamForUser(paramsList, userId, teamId, isOrSearch, includeDeletedChannels, page, perPage)

}

func (s *RetryLayerPostStore) Update(newPost *model.Post, oldPost *model.Post) (*model.Post, error) {

	tries := 0
	for {
		result, err := s.PostStore.Update(newPost, oldPost)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {

	tries := 0
	for {
		result, err := s.PreferenceStore.CleanupFlagsBatch(limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPreferenceStore) Delete(userId string, category string, name string) error {

	tries := 0
	for {
		err := s.PreferenceStore.Delete(userId, category, name)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPreferenceStore) DeleteCategory(userId string, category string) error {

	tries := 0
	for {
		err := s.PreferenceStore.DeleteCategory(userId, category)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPreferenceStore) DeleteCategoryAndName(category string, name string) error {

	tries := 0
	for {
		err := s.PreferenceStore.DeleteCategoryAndName(category, name)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPreferenceStore) Get(userId string, category string, name string) (*model.Preference, error) {

	tries := 0
	for {
		result, err := s.PreferenceStore.Get(userId, category, name)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPreferenceStore) GetAll(userId string) (model.Preferences, error) {

	tries := 0
	for {
		result, err := s.PreferenceStore.GetAll(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPreferenceStore) GetCategory(userId string, category string) (model.Preferences, error) {

	tries := 0
	for {
		result, err := s.PreferenceStore.GetCategory(userId, category)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerPreferenceStore) PermanentDeleteByUser(userId string) error {

	tries := 0
	for {
		err := s.PreferenceStore.PermanentDeleteByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerPreferenceStore) Save(preferences *model.Preferences) error {

	tries := 0
	for {
		err := s.PreferenceStore.Save(preferences)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, error) {

	tries := 0
	for {
		result, err := s.ReactionStore.BulkGetForPosts(postIds)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, error) {

	tries := 0
	for {
		result, err := s.ReactionStore.Delete(reaction)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerReactionStore) DeleteAllWithEmojiName(emojiName string) error {

	tries := 0
	for {
		err := s.ReactionStore.DeleteAllWithEmojiName(emojiName)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, error) {

	tries := 0
	for {
		result, err := s.ReactionStore.GetForPost(postId, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {

	tries := 0
	for {
		result, err := s.ReactionStore.PermanentDeleteBatch(endTime, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerReactionStore) Save(reaction *model.Reaction) (*model.Reaction, error) {

	tries := 0
	for {
		result, err := s.ReactionStore.Save(reaction)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) AllChannelSchemeRoles() ([]*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.AllChannelSchemeRoles()
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) ChannelHigherScopedPermissions(roleNames []string) (map[string]*model.RolePermissions, error) {

	tries := 0
	for {
		result, err := s.RoleStore.ChannelHigherScopedPermissions(roleNames)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) ChannelRolesUnderTeamRole(roleName string) ([]*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.ChannelRolesUnderTeamRole(roleName)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) Delete(roleId string) (*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.Delete(roleId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) Get(roleId string) (*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.Get(roleId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) GetAll() ([]*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.GetAll()
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) GetByName(name string) (*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.GetByName(name)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) GetByNames(names []string) ([]*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.GetByNames(names)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerRoleStore) PermanentDeleteAll() error {

	tries := 0
	for {
		err := s.RoleStore.PermanentDeleteAll()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerRoleStore) Save(role *model.Role) (*model.Role, error) {

	tries := 0
	for {
		result, err := s.RoleStore.Save(role)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) CountByScope(scope string) (int64, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.CountByScope(scope)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) CountWithoutPermission(scope string, permissionID string, roleScope model.RoleScope, roleType model.RoleType) (int64, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.CountWithoutPermission(scope, permissionID, roleScope, roleType)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) Delete(schemeId string) (*model.Scheme, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.Delete(schemeId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) Get(schemeId string) (*model.Scheme, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.Get(schemeId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) GetAllPage(scope string, offset int, limit int) ([]*model.Scheme, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.GetAllPage(scope, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) GetByName(schemeName string) (*model.Scheme, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.GetByName(schemeName)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSchemeStore) PermanentDeleteAll() error {

	tries := 0
	for {
		err := s.SchemeStore.PermanentDeleteAll()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSchemeStore) Save(scheme *model.Scheme) (*model.Scheme, error) {

	tries := 0
	for {
		result, err := s.SchemeStore.Save(scheme)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) AnalyticsSessionCount() (int64, error) {

	tries := 0
	for {
		result, err := s.SessionStore.AnalyticsSessionCount()
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) Cleanup(expiryTime int64, batchSize int64) {

	s.SessionStore.Cleanup(expiryTime, batchSize)

}

func (s *RetryLayerSessionStore) Get(sessionIdOrToken string) (*model.Session, error) {

	tries := 0
	for {
		result, err := s.SessionStore.Get(sessionIdOrToken)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) GetSessions(userId string) ([]*model.Session, error) {

	tries := 0
	for {
		result, err := s.SessionStore.GetSessions(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) GetSessionsExpired(thresholdMillis int64, mobileOnly bool, unnotifiedOnly bool) ([]*model.Session, error) {

	tries := 0
	for {
		result, err := s.SessionStore.GetSessionsExpired(thresholdMillis, mobileOnly, unnotifiedOnly)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) GetSessionsWithActiveDeviceIds(userId string) ([]*model.Session, error) {

	tries := 0
	for {
		result, err := s.SessionStore.GetSessionsWithActiveDeviceIds(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) PermanentDeleteSessionsByUser(teamId string) error {

	tries := 0
	for {
		err := s.SessionStore.PermanentDeleteSessionsByUser(teamId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) Remove(sessionIdOrToken string) error {

	tries := 0
	for {
		err := s.SessionStore.Remove(sessionIdOrToken)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) RemoveAllSessions() error {

	tries := 0
	for {
		err := s.SessionStore.RemoveAllSessions()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) Save(session *model.Session) (*model.Session, error) {

	tries := 0
	for {
		result, err := s.SessionStore.Save(session)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) (string, error) {

	tries := 0
	for {
		result, err := s.SessionStore.UpdateDeviceId(id, deviceId, expiresAt)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateExpiredNotify(sessionid string, notified bool) error {

	tries := 0
	for {
		err := s.SessionStore.UpdateExpiredNotify(sessionid, notified)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateExpiresAt(sessionId string, time int64) error {

	tries := 0
	for {
		err := s.SessionStore.UpdateExpiresAt(sessionId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateLastActivityAt(sessionId string, time int64) error {

	tries := 0
	for {
		err := s.SessionStore.UpdateLastActivityAt(sessionId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateProps(session *model.Session) error {

	tries := 0
	for {
		err := s.SessionStore.UpdateProps(session)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSessionStore) UpdateRoles(userId string, roles string) (string, error) {

	tries := 0
	for {
		result, err := s.SessionStore.UpdateRoles(userId, roles)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerStatusStore) Get(userId string) (*model.Status, error) {

	tries := 0
	for {
		result, err := s.StatusStore.Get(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerStatusStore) GetByIds(userIds []string) ([]*model.Status, error) {

	tries := 0
	for {
		result, err := s.StatusStore.GetByIds(userIds)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerStatusStore) GetTotalActiveUsersCount() (int64, error) {

	tries := 0
	for {
		result, err := s.StatusStore.GetTotalActiveUsersCount()
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerStatusStore) ResetAll() error {

	tries := 0
	for {
		err := s.StatusStore.ResetAll()
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerStatusStore) SaveOrUpdate(status *model.Status) error {

	tries := 0
	for {
		err := s.StatusStore.SaveOrUpdate(status)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) error {

	tries := 0
	for {
		err := s.StatusStore.UpdateLastActivityAt(userId, lastActivityAt)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSystemStore) Get() (model.StringMap, error) {

	tries := 0
	for {
		result, err := s.SystemStore.Get()
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSystemStore) GetByName(name string) (*model.System, error) {

	tries := 0
	for {
		result, err := s.SystemStore.GetByName(name)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSystemStore) InsertIfExists(system *model.System) (*model.System, error) {

	tries := 0
	for {
		result, err := s.SystemStore.InsertIfExists(system)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSystemStore) PermanentDeleteByName(name string) (*model.System, error) {

	tries := 0
	for {
		result, err := s.SystemStore.PermanentDeleteByName(name)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerSystemStore) Save(system *model.System) error {

	tries := 0
	for {
		err := s.SystemStore.Save(system)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSystemStore) SaveOrUpdate(system *model.System) error {

	tries := 0
	for {
		err := s.SystemStore.SaveOrUpdate(system)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerSystemStore) Update(system *model.System) error {

	tries := 0
	for {
		err := s.SystemStore.Update(system)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) (int64, *model.AppError) {

	return s.TeamStore.AnalyticsGetTeamCountForScheme(schemeId)

}

func (s *RetryLayerTeamStore) AnalyticsPrivateTeamCount() (int64, *model.AppError) {

	return s.TeamStore.AnalyticsPrivateTeamCount()

}

func (s *RetryLayerTeamStore) AnalyticsPublicTeamCount() (int64, *model.AppError) {

	return s.TeamStore.AnalyticsPublicTeamCount()

}

func (s *RetryLayerTeamStore) AnalyticsTeamCount(includeDeleted bool) (int64, *model.AppError) {

	return s.TeamStore.AnalyticsTeamCount(includeDeleted)

}

func (s *RetryLayerTeamStore) ClearAllCustomRoleAssignments() *model.AppError {

	return s.TeamStore.ClearAllCustomRoleAssignments()

}

func (s *RetryLayerTeamStore) ClearCaches() {

	s.TeamStore.ClearCaches()

}

func (s *RetryLayerTeamStore) Get(id string) (*model.Team, *model.AppError) {

	return s.TeamStore.Get(id)

}

func (s *RetryLayerTeamStore) GetActiveMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {

	return s.TeamStore.GetActiveMemberCount(teamId, restrictions)

}

func (s *RetryLayerTeamStore) GetAll() ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAll()

}

func (s *RetryLayerTeamStore) GetAllForExportAfter(limit int, afterId string) ([]*model.TeamForExport, *model.AppError) {

	return s.TeamStore.GetAllForExportAfter(limit, afterId)

}

func (s *RetryLayerTeamStore) GetAllPage(offset int, limit int) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllPage(offset, limit)

}

func (s *RetryLayerTeamStore) GetAllPrivateTeamListing() ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllPrivateTeamListing()

}

func (s *RetryLayerTeamStore) GetAllPrivateTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllPrivateTeamPageListing(offset, limit)

}

func (s *RetryLayerTeamStore) GetAllPublicTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllPublicTeamPageListing(offset, limit)

}

func (s *RetryLayerTeamStore) GetAllTeamListing() ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllTeamListing()

}

func (s *RetryLayerTeamStore) GetAllTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetAllTeamPageListing(offset, limit)

}

func (s *RetryLayerTeamStore) GetByInviteId(inviteId string) (*model.Team, *model.AppError) {

	return s.TeamStore.GetByInviteId(inviteId)

}

func (s *RetryLayerTeamStore) GetByName(name string) (*model.Team, *model.AppError) {

	return s.TeamStore.GetByName(name)

}

func (s *RetryLayerTeamStore) GetByNames(name []string) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetByNames(name)

}

func (s *RetryLayerTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {

	return s.TeamStore.GetChannelUnreadsForAllTeams(excludeTeamId, userId)

}

func (s *RetryLayerTeamStore) GetChannelUnreadsForTeam(teamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {

	return s.TeamStore.GetChannelUnreadsForTeam(teamId, userId)

}

func (s *RetryLayerTeamStore) GetMember(teamId string, userId string) (*model.TeamMember, *model.AppError) {

	return s.TeamStore.GetMember(teamId, userId)

}

func (s *RetryLayerTeamStore) GetMembers(teamId string, offset int, limit int, teamMembersGetOptions *model.TeamMembersGetOptions) ([]*model.TeamMember, *model.AppError) {

	return s.TeamStore.GetMembers(teamId, offset, limit, teamMembersGetOptions)

}

func (s *RetryLayerTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) ([]*model.TeamMember, *model.AppError) {

	return s.TeamStore.GetMembersByIds(teamId, userIds, restrictions)

}

func (s *RetryLayerTeamStore) GetTeamMembersForExport(userId string) ([]*model.TeamMemberForExport, *model.AppError) {

	return s.TeamStore.GetTeamMembersForExport(userId)

}

func (s *RetryLayerTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetTeamsByScheme(schemeId, offset, limit)

}

func (s *RetryLayerTeamStore) GetTeamsByUserId(userId string) ([]*model.Team, *model.AppError) {

	return s.TeamStore.GetTeamsByUserId(userId)

}

func (s *RetryLayerTeamStore) GetTeamsForUser(userId string) ([]*model.TeamMember, *model.AppError) {

	return s.TeamStore.GetTeamsForUser(userId)

}

func (s *RetryLayerTeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {

	return s.TeamStore.GetTeamsForUserWithPagination(userId, page, perPage)

}

func (s *RetryLayerTeamStore) GetTotalMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {

	return s.TeamStore.GetTotalMemberCount(teamId, restrictions)

}

func (s *RetryLayerTeamStore) GetUserTeamIds(userId string, allowFromCache bool) ([]string, *model.AppError) {

	return s.TeamStore.GetUserTeamIds(userId, allowFromCache)

}

func (s *RetryLayerTeamStore) GroupSyncedTeamCount() (int64, *model.AppError) {

	return s.TeamStore.GroupSyncedTeamCount()

}

func (s *RetryLayerTeamStore) InvalidateAllTeamIdsForUser(userId string) {

	s.TeamStore.InvalidateAllTeamIdsForUser(userId)

}

func (s *RetryLayerTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) (map[string]string, *model.AppError) {

	return s.TeamStore.MigrateTeamMembers(fromTeamId, fromUserId)

}

func (s *RetryLayerTeamStore) PermanentDelete(teamId string) *model.AppError {

	return s.TeamStore.PermanentDelete(teamId)

}

func (s *RetryLayerTeamStore) RemoveAllMembersByTeam(teamId string) *model.AppError {

	return s.TeamStore.RemoveAllMembersByTeam(teamId)

}

func (s *RetryLayerTeamStore) RemoveAllMembersByUser(userId string) *model.AppError {

	return s.TeamStore.RemoveAllMembersByUser(userId)

}

func (s *RetryLayerTeamStore) RemoveMember(teamId string, userId string) *model.AppError {

	return s.TeamStore.RemoveMember(teamId, userId)

}

func (s *RetryLayerTeamStore) RemoveMembers(teamId string, userIds []string) *model.AppError {

	return s.TeamStore.RemoveMembers(teamId, userIds)

}

func (s *RetryLayerTeamStore) ResetAllTeamSchemes() *model.AppError {

	return s.TeamStore.ResetAllTeamSchemes()

}

func (s *RetryLayerTeamStore) Save(team *model.Team) (*model.Team, *model.AppError) {

	return s.TeamStore.Save(team)

}

func (s *RetryLayerTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) (*model.TeamMember, error) {

	tries := 0
	for {
		result, err := s.TeamStore.SaveMember(member, maxUsersPerTeam)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTeamStore) SaveMultipleMembers(members []*model.TeamMember, maxUsersPerTeam int) ([]*model.TeamMember, error) {

	tries := 0
	for {
		result, err := s.TeamStore.SaveMultipleMembers(members, maxUsersPerTeam)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTeamStore) SearchAll(term string, opts *model.TeamSearch) ([]*model.Team, *model.AppError) {

	return s.TeamStore.SearchAll(term, opts)

}

func (s *RetryLayerTeamStore) SearchAllPaged(term string, opts *model.TeamSearch) ([]*model.Team, int64, *model.AppError) {

	return s.TeamStore.SearchAllPaged(term, opts)

}

func (s *RetryLayerTeamStore) SearchOpen(term string) ([]*model.Team, *model.AppError) {

	return s.TeamStore.SearchOpen(term)

}

func (s *RetryLayerTeamStore) SearchPrivate(term string) ([]*model.Team, *model.AppError) {

	return s.TeamStore.SearchPrivate(term)

}

func (s *RetryLayerTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {

	return s.TeamStore.Update(team)

}

func (s *RetryLayerTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) *model.AppError {

	return s.TeamStore.UpdateLastTeamIconUpdate(teamId, curTime)

}

func (s *RetryLayerTeamStore) UpdateMember(member *model.TeamMember) (*model.TeamMember, *model.AppError) {

	return s.TeamStore.UpdateMember(member)

}

func (s *RetryLayerTeamStore) UpdateMembersRole(teamID string, userIDs []string) *model.AppError {

	return s.TeamStore.UpdateMembersRole(teamID, userIDs)

}

func (s *RetryLayerTeamStore) UpdateMultipleMembers(members []*model.TeamMember) ([]*model.TeamMember, *model.AppError) {

	return s.TeamStore.UpdateMultipleMembers(members)

}

func (s *RetryLayerTeamStore) UserBelongsToTeams(userId string, teamIds []string) (bool, *model.AppError) {

	return s.TeamStore.UserBelongsToTeams(userId, teamIds)

}

func (s *RetryLayerTermsOfServiceStore) Get(id string, allowFromCache bool) (*model.TermsOfService, error) {

	tries := 0
	for {
		result, err := s.TermsOfServiceStore.Get(id, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTermsOfServiceStore) GetLatest(allowFromCache bool) (*model.TermsOfService, error) {

	tries := 0
	for {
		result, err := s.TermsOfServiceStore.GetLatest(allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTermsOfServiceStore) Save(termsOfService *model.TermsOfService) (*model.TermsOfService, error) {

	tries := 0
	for {
		result, err := s.TermsOfServiceStore.Save(termsOfService)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTokenStore) Cleanup() {

	s.TokenStore.Cleanup()

}

func (s *RetryLayerTokenStore) Delete(token string) error {

	tries := 0
	for {
		err := s.TokenStore.Delete(token)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerTokenStore) GetByToken(token string) (*model.Token, error) {

	tries := 0
	for {
		result, err := s.TokenStore.GetByToken(token)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerTokenStore) RemoveAllTokensByType(tokenType string) error {

	tries := 0
	for {
		err := s.TokenStore.RemoveAllTokensByType(tokenType)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerTokenStore) Save(recovery *model.Token) error {

	tries := 0
	for {
		err := s.TokenStore.Save(recovery)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserStore) AnalyticsActiveCount(time int64, options model.UserCountOptions) (int64, *model.AppError) {

	return s.UserStore.AnalyticsActiveCount(time, options)

}

func (s *RetryLayerUserStore) AnalyticsGetGuestCount() (int64, *model.AppError) {

	return s.UserStore.AnalyticsGetGuestCount()

}

func (s *RetryLayerUserStore) AnalyticsGetInactiveUsersCount() (int64, *model.AppError) {

	return s.UserStore.AnalyticsGetInactiveUsersCount()

}

func (s *RetryLayerUserStore) AnalyticsGetSystemAdminCount() (int64, *model.AppError) {

	return s.UserStore.AnalyticsGetSystemAdminCount()

}

func (s *RetryLayerUserStore) AutocompleteUsersInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) (*model.UserAutocompleteInChannel, *model.AppError) {

	return s.UserStore.AutocompleteUsersInChannel(teamId, channelId, term, options)

}

func (s *RetryLayerUserStore) ClearAllCustomRoleAssignments() *model.AppError {

	return s.UserStore.ClearAllCustomRoleAssignments()

}

func (s *RetryLayerUserStore) ClearCaches() {

	s.UserStore.ClearCaches()

}

func (s *RetryLayerUserStore) Count(options model.UserCountOptions) (int64, *model.AppError) {

	return s.UserStore.Count(options)

}

func (s *RetryLayerUserStore) DeactivateGuests() ([]string, *model.AppError) {

	return s.UserStore.DeactivateGuests()

}

func (s *RetryLayerUserStore) DemoteUserToGuest(userID string) *model.AppError {

	return s.UserStore.DemoteUserToGuest(userID)

}

func (s *RetryLayerUserStore) Get(id string) (*model.User, *model.AppError) {

	return s.UserStore.Get(id)

}

func (s *RetryLayerUserStore) GetAll() ([]*model.User, *model.AppError) {

	return s.UserStore.GetAll()

}

func (s *RetryLayerUserStore) GetAllAfter(limit int, afterId string) ([]*model.User, *model.AppError) {

	return s.UserStore.GetAllAfter(limit, afterId)

}

func (s *RetryLayerUserStore) GetAllNotInAuthService(authServices []string) ([]*model.User, *model.AppError) {

	return s.UserStore.GetAllNotInAuthService(authServices)

}

func (s *RetryLayerUserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetAllProfiles(options)

}

func (s *RetryLayerUserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) (map[string]*model.User, *model.AppError) {

	return s.UserStore.GetAllProfilesInChannel(channelId, allowFromCache)

}

func (s *RetryLayerUserStore) GetAllUsingAuthService(authService string) ([]*model.User, *model.AppError) {

	return s.UserStore.GetAllUsingAuthService(authService)

}

func (s *RetryLayerUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) (int64, *model.AppError) {

	return s.UserStore.GetAnyUnreadPostCountForChannel(userId, channelId)

}

func (s *RetryLayerUserStore) GetByAuth(authData *string, authService string) (*model.User, *model.AppError) {

	return s.UserStore.GetByAuth(authData, authService)

}

func (s *RetryLayerUserStore) GetByEmail(email string) (*model.User, *model.AppError) {

	return s.UserStore.GetByEmail(email)

}

func (s *RetryLayerUserStore) GetByUsername(username string) (*model.User, *model.AppError) {

	return s.UserStore.GetByUsername(username)

}

func (s *RetryLayerUserStore) GetChannelGroupUsers(channelID string) ([]*model.User, *model.AppError) {

	return s.UserStore.GetChannelGroupUsers(channelID)

}

func (s *RetryLayerUserStore) GetEtagForAllProfiles() string {

	return s.UserStore.GetEtagForAllProfiles()

}

func (s *RetryLayerUserStore) GetEtagForProfiles(teamId string) string {

	return s.UserStore.GetEtagForProfiles(teamId)

}

func (s *RetryLayerUserStore) GetEtagForProfilesNotInTeam(teamId string) string {

	return s.UserStore.GetEtagForProfilesNotInTeam(teamId)

}

func (s *RetryLayerUserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) (*model.User, *model.AppError) {

	return s.UserStore.GetForLogin(loginId, allowSignInWithUsername, allowSignInWithEmail)

}

func (s *RetryLayerUserStore) GetKnownUsers(userID string) ([]string, *model.AppError) {

	return s.UserStore.GetKnownUsers(userID)

}

func (s *RetryLayerUserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetNewUsersForTeam(teamId, offset, limit, viewRestrictions)

}

func (s *RetryLayerUserStore) GetProfileByGroupChannelIdsForUser(userId string, channelIds []string) (map[string][]*model.User, *model.AppError) {

	return s.UserStore.GetProfileByGroupChannelIdsForUser(userId, channelIds)

}

func (s *RetryLayerUserStore) GetProfileByIds(userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfileByIds(userIds, options, allowFromCache)

}

func (s *RetryLayerUserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfiles(options)

}

func (s *RetryLayerUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesByUsernames(usernames, viewRestrictions)

}

func (s *RetryLayerUserStore) GetProfilesInChannel(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesInChannel(channelId, offset, limit)

}

func (s *RetryLayerUserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesInChannelByStatus(channelId, offset, limit)

}

func (s *RetryLayerUserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesNotInChannel(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)

}

func (s *RetryLayerUserStore) GetProfilesNotInTeam(teamId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesNotInTeam(teamId, groupConstrained, offset, limit, viewRestrictions)

}

func (s *RetryLayerUserStore) GetProfilesWithoutTeam(options *model.UserGetOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetProfilesWithoutTeam(options)

}

func (s *RetryLayerUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {

	return s.UserStore.GetRecentlyActiveUsersForTeam(teamId, offset, limit, viewRestrictions)

}

func (s *RetryLayerUserStore) GetSystemAdminProfiles() (map[string]*model.User, *model.AppError) {

	return s.UserStore.GetSystemAdminProfiles()

}

func (s *RetryLayerUserStore) GetTeamGroupUsers(teamID string) ([]*model.User, *model.AppError) {

	return s.UserStore.GetTeamGroupUsers(teamID)

}

func (s *RetryLayerUserStore) GetUnreadCount(userId string) (int64, *model.AppError) {

	return s.UserStore.GetUnreadCount(userId)

}

func (s *RetryLayerUserStore) GetUnreadCountForChannel(userId string, channelId string) (int64, *model.AppError) {

	return s.UserStore.GetUnreadCountForChannel(userId, channelId)

}

func (s *RetryLayerUserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.UserForIndexing, *model.AppError) {

	return s.UserStore.GetUsersBatchForIndexing(startTime, endTime, limit)

}

func (s *RetryLayerUserStore) InferSystemInstallDate() (int64, *model.AppError) {

	return s.UserStore.InferSystemInstallDate()

}

func (s *RetryLayerUserStore) InvalidateProfileCacheForUser(userId string) {

	s.UserStore.InvalidateProfileCacheForUser(userId)

}

func (s *RetryLayerUserStore) InvalidateProfilesInChannelCache(channelId string) {

	s.UserStore.InvalidateProfilesInChannelCache(channelId)

}

func (s *RetryLayerUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {

	s.UserStore.InvalidateProfilesInChannelCacheByUser(userId)

}

func (s *RetryLayerUserStore) PermanentDelete(userId string) *model.AppError {

	return s.UserStore.PermanentDelete(userId)

}

func (s *RetryLayerUserStore) PromoteGuestToUser(userID string) *model.AppError {

	return s.UserStore.PromoteGuestToUser(userID)

}

func (s *RetryLayerUserStore) ResetLastPictureUpdate(userId string) *model.AppError {

	return s.UserStore.ResetLastPictureUpdate(userId)

}

func (s *RetryLayerUserStore) Save(user *model.User) (*model.User, *model.AppError) {

	return s.UserStore.Save(user)

}

func (s *RetryLayerUserStore) Search(teamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.Search(teamId, term, options)

}

func (s *RetryLayerUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.SearchInChannel(channelId, term, options)

}

func (s *RetryLayerUserStore) SearchInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.SearchInGroup(groupID, term, options)

}

func (s *RetryLayerUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.SearchNotInChannel(teamId, channelId, term, options)

}

func (s *RetryLayerUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.SearchNotInTeam(notInTeamId, term, options)

}

func (s *RetryLayerUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {

	return s.UserStore.SearchWithoutTeam(term, options)

}

func (s *RetryLayerUserStore) Update(user *model.User, allowRoleUpdate bool) (*model.UserUpdate, *model.AppError) {

	return s.UserStore.Update(user, allowRoleUpdate)

}

func (s *RetryLayerUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) (string, *model.AppError) {

	return s.UserStore.UpdateAuthData(userId, service, authData, email, resetMfa)

}

func (s *RetryLayerUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) *model.AppError {

	return s.UserStore.UpdateFailedPasswordAttempts(userId, attempts)

}

func (s *RetryLayerUserStore) UpdateLastPictureUpdate(userId string) *model.AppError {

	return s.UserStore.UpdateLastPictureUpdate(userId)

}

func (s *RetryLayerUserStore) UpdateMfaActive(userId string, active bool) *model.AppError {

	return s.UserStore.UpdateMfaActive(userId, active)

}

func (s *RetryLayerUserStore) UpdateMfaSecret(userId string, secret string) *model.AppError {

	return s.UserStore.UpdateMfaSecret(userId, secret)

}

func (s *RetryLayerUserStore) UpdatePassword(userId string, newPassword string) *model.AppError {

	return s.UserStore.UpdatePassword(userId, newPassword)

}

func (s *RetryLayerUserStore) UpdateUpdateAt(userId string) (int64, *model.AppError) {

	return s.UserStore.UpdateUpdateAt(userId)

}

func (s *RetryLayerUserStore) VerifyEmail(userId string, email string) (string, *model.AppError) {

	return s.UserStore.VerifyEmail(userId, email)

}

func (s *RetryLayerUserAccessTokenStore) Delete(tokenId string) error {

	tries := 0
	for {
		err := s.UserAccessTokenStore.Delete(tokenId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) DeleteAllForUser(userId string) error {

	tries := 0
	for {
		err := s.UserAccessTokenStore.DeleteAllForUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) Get(tokenId string) (*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.Get(tokenId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) GetAll(offset int, limit int) ([]*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.GetAll(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) GetByToken(tokenString string) (*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.GetByToken(tokenString)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) GetByUser(userId string, page int, perPage int) ([]*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.GetByUser(userId, page, perPage)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) Save(token *model.UserAccessToken) (*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.Save(token)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) Search(term string) ([]*model.UserAccessToken, error) {

	tries := 0
	for {
		result, err := s.UserAccessTokenStore.Search(term)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) UpdateTokenDisable(tokenId string) error {

	tries := 0
	for {
		err := s.UserAccessTokenStore.UpdateTokenDisable(tokenId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserAccessTokenStore) UpdateTokenEnable(tokenId string) error {

	tries := 0
	for {
		err := s.UserAccessTokenStore.UpdateTokenEnable(tokenId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) error {

	tries := 0
	for {
		err := s.UserTermsOfServiceStore.Delete(userId, termsOfServiceId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerUserTermsOfServiceStore) GetByUser(userId string) (*model.UserTermsOfService, error) {

	tries := 0
	for {
		result, err := s.UserTermsOfServiceStore.GetByUser(userId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerUserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) (*model.UserTermsOfService, error) {

	tries := 0
	for {
		result, err := s.UserTermsOfServiceStore.Save(userTermsOfService)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) AnalyticsIncomingCount(teamId string) (int64, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.AnalyticsIncomingCount(teamId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.AnalyticsOutgoingCount(teamId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) ClearCaches() {

	s.WebhookStore.ClearCaches()

}

func (s *RetryLayerWebhookStore) DeleteIncoming(webhookId string, time int64) error {

	tries := 0
	for {
		err := s.WebhookStore.DeleteIncoming(webhookId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) DeleteOutgoing(webhookId string, time int64) error {

	tries := 0
	for {
		err := s.WebhookStore.DeleteOutgoing(webhookId, time)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncoming(id, allowFromCache)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncomingByChannel(channelId)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncomingByTeam(teamId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncomingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncomingByTeamByUser(teamId, userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncomingList(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetIncomingListByUser(userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetIncomingListByUser(userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoing(id)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingByChannel(channelId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingByChannelByUser(channelId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingByChannelByUser(channelId, userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingByTeam(teamId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingByTeamByUser(teamId, userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingList(offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) GetOutgoingListByUser(userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.GetOutgoingListByUser(userId, offset, limit)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) InvalidateWebhookCache(webhook string) {

	s.WebhookStore.InvalidateWebhookCache(webhook)

}

func (s *RetryLayerWebhookStore) PermanentDeleteIncomingByChannel(channelId string) error {

	tries := 0
	for {
		err := s.WebhookStore.PermanentDeleteIncomingByChannel(channelId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) PermanentDeleteIncomingByUser(userId string) error {

	tries := 0
	for {
		err := s.WebhookStore.PermanentDeleteIncomingByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) error {

	tries := 0
	for {
		err := s.WebhookStore.PermanentDeleteOutgoingByChannel(channelId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) PermanentDeleteOutgoingByUser(userId string) error {

	tries := 0
	for {
		err := s.WebhookStore.PermanentDeleteOutgoingByUser(userId)
		if err == nil {
			return nil
		}
		if !isRepeatableError(err) {
			return err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return err
		}
	}

}

func (s *RetryLayerWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.SaveIncoming(webhook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.SaveOutgoing(webhook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.UpdateIncoming(webhook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayerWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {

	tries := 0
	for {
		result, err := s.WebhookStore.UpdateOutgoing(hook)
		if err == nil {
			return result, nil
		}
		if !isRepeatableError(err) {
			return result, err
		}
		tries++
		if tries >= 3 {
			err = errors.Wrap(err, "giving up after 3 consecutive repeatable transaction failures")
			return result, err
		}
	}

}

func (s *RetryLayer) Close() {
	s.Store.Close()
}

func (s *RetryLayer) DropAllTables() {
	s.Store.DropAllTables()
}

func (s *RetryLayer) GetCurrentSchemaVersion() string {
	return s.Store.GetCurrentSchemaVersion()
}

func (s *RetryLayer) LockToMaster() {
	s.Store.LockToMaster()
}

func (s *RetryLayer) MarkSystemRanUnitTests() {
	s.Store.MarkSystemRanUnitTests()
}

func (s *RetryLayer) SetContext(context context.Context) {
	s.Store.SetContext(context)
}

func (s *RetryLayer) TotalMasterDbConnections() int {
	return s.Store.TotalMasterDbConnections()
}

func (s *RetryLayer) TotalReadDbConnections() int {
	return s.Store.TotalReadDbConnections()
}

func (s *RetryLayer) TotalSearchDbConnections() int {
	return s.Store.TotalSearchDbConnections()
}

func (s *RetryLayer) UnlockFromMaster() {
	s.Store.UnlockFromMaster()
}

func New(childStore store.Store) *RetryLayer {
	newStore := RetryLayer{
		Store: childStore,
	}

	newStore.AuditStore = &RetryLayerAuditStore{AuditStore: childStore.Audit(), Root: &newStore}
	newStore.BotStore = &RetryLayerBotStore{BotStore: childStore.Bot(), Root: &newStore}
	newStore.ChannelStore = &RetryLayerChannelStore{ChannelStore: childStore.Channel(), Root: &newStore}
	newStore.ChannelMemberHistoryStore = &RetryLayerChannelMemberHistoryStore{ChannelMemberHistoryStore: childStore.ChannelMemberHistory(), Root: &newStore}
	newStore.ClusterDiscoveryStore = &RetryLayerClusterDiscoveryStore{ClusterDiscoveryStore: childStore.ClusterDiscovery(), Root: &newStore}
	newStore.CommandStore = &RetryLayerCommandStore{CommandStore: childStore.Command(), Root: &newStore}
	newStore.CommandWebhookStore = &RetryLayerCommandWebhookStore{CommandWebhookStore: childStore.CommandWebhook(), Root: &newStore}
	newStore.ComplianceStore = &RetryLayerComplianceStore{ComplianceStore: childStore.Compliance(), Root: &newStore}
	newStore.EmojiStore = &RetryLayerEmojiStore{EmojiStore: childStore.Emoji(), Root: &newStore}
	newStore.FileInfoStore = &RetryLayerFileInfoStore{FileInfoStore: childStore.FileInfo(), Root: &newStore}
	newStore.GroupStore = &RetryLayerGroupStore{GroupStore: childStore.Group(), Root: &newStore}
	newStore.JobStore = &RetryLayerJobStore{JobStore: childStore.Job(), Root: &newStore}
	newStore.LicenseStore = &RetryLayerLicenseStore{LicenseStore: childStore.License(), Root: &newStore}
	newStore.LinkMetadataStore = &RetryLayerLinkMetadataStore{LinkMetadataStore: childStore.LinkMetadata(), Root: &newStore}
	newStore.OAuthStore = &RetryLayerOAuthStore{OAuthStore: childStore.OAuth(), Root: &newStore}
	newStore.PluginStore = &RetryLayerPluginStore{PluginStore: childStore.Plugin(), Root: &newStore}
	newStore.PostStore = &RetryLayerPostStore{PostStore: childStore.Post(), Root: &newStore}
	newStore.PreferenceStore = &RetryLayerPreferenceStore{PreferenceStore: childStore.Preference(), Root: &newStore}
	newStore.ReactionStore = &RetryLayerReactionStore{ReactionStore: childStore.Reaction(), Root: &newStore}
	newStore.RoleStore = &RetryLayerRoleStore{RoleStore: childStore.Role(), Root: &newStore}
	newStore.SchemeStore = &RetryLayerSchemeStore{SchemeStore: childStore.Scheme(), Root: &newStore}
	newStore.SessionStore = &RetryLayerSessionStore{SessionStore: childStore.Session(), Root: &newStore}
	newStore.StatusStore = &RetryLayerStatusStore{StatusStore: childStore.Status(), Root: &newStore}
	newStore.SystemStore = &RetryLayerSystemStore{SystemStore: childStore.System(), Root: &newStore}
	newStore.TeamStore = &RetryLayerTeamStore{TeamStore: childStore.Team(), Root: &newStore}
	newStore.TermsOfServiceStore = &RetryLayerTermsOfServiceStore{TermsOfServiceStore: childStore.TermsOfService(), Root: &newStore}
	newStore.TokenStore = &RetryLayerTokenStore{TokenStore: childStore.Token(), Root: &newStore}
	newStore.UserStore = &RetryLayerUserStore{UserStore: childStore.User(), Root: &newStore}
	newStore.UserAccessTokenStore = &RetryLayerUserAccessTokenStore{UserAccessTokenStore: childStore.UserAccessToken(), Root: &newStore}
	newStore.UserTermsOfServiceStore = &RetryLayerUserTermsOfServiceStore{UserTermsOfServiceStore: childStore.UserTermsOfService(), Root: &newStore}
	newStore.WebhookStore = &RetryLayerWebhookStore{WebhookStore: childStore.Webhook(), Root: &newStore}
	return &newStore
}
