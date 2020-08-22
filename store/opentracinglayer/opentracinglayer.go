// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package opentracinglayer

import (
	"context"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/services/tracing"
	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/opentracing/opentracing-go/ext"
	spanlog "github.com/opentracing/opentracing-go/log"
)

type OpenTracingLayer struct {
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

func (s *OpenTracingLayer) Audit() store.AuditStore {
	return s.AuditStore
}

func (s *OpenTracingLayer) Bot() store.BotStore {
	return s.BotStore
}

func (s *OpenTracingLayer) Channel() store.ChannelStore {
	return s.ChannelStore
}

func (s *OpenTracingLayer) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.ChannelMemberHistoryStore
}

func (s *OpenTracingLayer) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.ClusterDiscoveryStore
}

func (s *OpenTracingLayer) Command() store.CommandStore {
	return s.CommandStore
}

func (s *OpenTracingLayer) CommandWebhook() store.CommandWebhookStore {
	return s.CommandWebhookStore
}

func (s *OpenTracingLayer) Compliance() store.ComplianceStore {
	return s.ComplianceStore
}

func (s *OpenTracingLayer) Emoji() store.EmojiStore {
	return s.EmojiStore
}

func (s *OpenTracingLayer) FileInfo() store.FileInfoStore {
	return s.FileInfoStore
}

func (s *OpenTracingLayer) Group() store.GroupStore {
	return s.GroupStore
}

func (s *OpenTracingLayer) Job() store.JobStore {
	return s.JobStore
}

func (s *OpenTracingLayer) License() store.LicenseStore {
	return s.LicenseStore
}

func (s *OpenTracingLayer) LinkMetadata() store.LinkMetadataStore {
	return s.LinkMetadataStore
}

func (s *OpenTracingLayer) OAuth() store.OAuthStore {
	return s.OAuthStore
}

func (s *OpenTracingLayer) Plugin() store.PluginStore {
	return s.PluginStore
}

func (s *OpenTracingLayer) Post() store.PostStore {
	return s.PostStore
}

func (s *OpenTracingLayer) Preference() store.PreferenceStore {
	return s.PreferenceStore
}

func (s *OpenTracingLayer) Reaction() store.ReactionStore {
	return s.ReactionStore
}

func (s *OpenTracingLayer) Role() store.RoleStore {
	return s.RoleStore
}

func (s *OpenTracingLayer) Scheme() store.SchemeStore {
	return s.SchemeStore
}

func (s *OpenTracingLayer) Session() store.SessionStore {
	return s.SessionStore
}

func (s *OpenTracingLayer) Status() store.StatusStore {
	return s.StatusStore
}

func (s *OpenTracingLayer) System() store.SystemStore {
	return s.SystemStore
}

func (s *OpenTracingLayer) Team() store.TeamStore {
	return s.TeamStore
}

func (s *OpenTracingLayer) TermsOfService() store.TermsOfServiceStore {
	return s.TermsOfServiceStore
}

func (s *OpenTracingLayer) Token() store.TokenStore {
	return s.TokenStore
}

func (s *OpenTracingLayer) User() store.UserStore {
	return s.UserStore
}

func (s *OpenTracingLayer) UserAccessToken() store.UserAccessTokenStore {
	return s.UserAccessTokenStore
}

func (s *OpenTracingLayer) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.UserTermsOfServiceStore
}

func (s *OpenTracingLayer) Webhook() store.WebhookStore {
	return s.WebhookStore
}

type OpenTracingLayerAuditStore struct {
	store.AuditStore
	Root *OpenTracingLayer
}

type OpenTracingLayerBotStore struct {
	store.BotStore
	Root *OpenTracingLayer
}

type OpenTracingLayerChannelStore struct {
	store.ChannelStore
	Root *OpenTracingLayer
}

type OpenTracingLayerChannelMemberHistoryStore struct {
	store.ChannelMemberHistoryStore
	Root *OpenTracingLayer
}

type OpenTracingLayerClusterDiscoveryStore struct {
	store.ClusterDiscoveryStore
	Root *OpenTracingLayer
}

type OpenTracingLayerCommandStore struct {
	store.CommandStore
	Root *OpenTracingLayer
}

type OpenTracingLayerCommandWebhookStore struct {
	store.CommandWebhookStore
	Root *OpenTracingLayer
}

type OpenTracingLayerComplianceStore struct {
	store.ComplianceStore
	Root *OpenTracingLayer
}

type OpenTracingLayerEmojiStore struct {
	store.EmojiStore
	Root *OpenTracingLayer
}

type OpenTracingLayerFileInfoStore struct {
	store.FileInfoStore
	Root *OpenTracingLayer
}

type OpenTracingLayerGroupStore struct {
	store.GroupStore
	Root *OpenTracingLayer
}

type OpenTracingLayerJobStore struct {
	store.JobStore
	Root *OpenTracingLayer
}

type OpenTracingLayerLicenseStore struct {
	store.LicenseStore
	Root *OpenTracingLayer
}

type OpenTracingLayerLinkMetadataStore struct {
	store.LinkMetadataStore
	Root *OpenTracingLayer
}

type OpenTracingLayerOAuthStore struct {
	store.OAuthStore
	Root *OpenTracingLayer
}

type OpenTracingLayerPluginStore struct {
	store.PluginStore
	Root *OpenTracingLayer
}

type OpenTracingLayerPostStore struct {
	store.PostStore
	Root *OpenTracingLayer
}

type OpenTracingLayerPreferenceStore struct {
	store.PreferenceStore
	Root *OpenTracingLayer
}

type OpenTracingLayerReactionStore struct {
	store.ReactionStore
	Root *OpenTracingLayer
}

type OpenTracingLayerRoleStore struct {
	store.RoleStore
	Root *OpenTracingLayer
}

type OpenTracingLayerSchemeStore struct {
	store.SchemeStore
	Root *OpenTracingLayer
}

type OpenTracingLayerSessionStore struct {
	store.SessionStore
	Root *OpenTracingLayer
}

type OpenTracingLayerStatusStore struct {
	store.StatusStore
	Root *OpenTracingLayer
}

type OpenTracingLayerSystemStore struct {
	store.SystemStore
	Root *OpenTracingLayer
}

type OpenTracingLayerTeamStore struct {
	store.TeamStore
	Root *OpenTracingLayer
}

type OpenTracingLayerTermsOfServiceStore struct {
	store.TermsOfServiceStore
	Root *OpenTracingLayer
}

type OpenTracingLayerTokenStore struct {
	store.TokenStore
	Root *OpenTracingLayer
}

type OpenTracingLayerUserStore struct {
	store.UserStore
	Root *OpenTracingLayer
}

type OpenTracingLayerUserAccessTokenStore struct {
	store.UserAccessTokenStore
	Root *OpenTracingLayer
}

type OpenTracingLayerUserTermsOfServiceStore struct {
	store.UserTermsOfServiceStore
	Root *OpenTracingLayer
}

type OpenTracingLayerWebhookStore struct {
	store.WebhookStore
	Root *OpenTracingLayer
}

func (s *OpenTracingLayerAuditStore) Get(user_id string, offset int, limit int) (model.Audits, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "AuditStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.AuditStore.Get(user_id, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerAuditStore) PermanentDeleteByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "AuditStore.PermanentDeleteByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.AuditStore.PermanentDeleteByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerAuditStore) Save(audit *model.Audit) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "AuditStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.AuditStore.Save(audit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerBotStore) Get(userId string, includeDeleted bool) (*model.Bot, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "BotStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.BotStore.Get(userId, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerBotStore) GetAll(options *model.BotGetOptions) ([]*model.Bot, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "BotStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.BotStore.GetAll(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerBotStore) PermanentDelete(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "BotStore.PermanentDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.BotStore.PermanentDelete(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerBotStore) Save(bot *model.Bot) (*model.Bot, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "BotStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.BotStore.Save(bot)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerBotStore) Update(bot *model.Bot) (*model.Bot, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "BotStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.BotStore.Update(bot)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.AnalyticsDeletedTypeCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.AnalyticsDeletedTypeCount(teamId, channelType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) AnalyticsTypeCount(teamId string, channelType string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.AnalyticsTypeCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.AnalyticsTypeCount(teamId, channelType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) AutocompleteInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.AutocompleteInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.AutocompleteInTeam(teamId, term, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.AutocompleteInTeamForSearch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.AutocompleteInTeamForSearch(teamId, userId, term, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) ClearAllCustomRoleAssignments() *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.ClearAllCustomRoleAssignments")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.ClearAllCustomRoleAssignments()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.ClearCaches()

}

func (s *OpenTracingLayerChannelStore) ClearSidebarOnTeamLeave(userId string, teamId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.ClearSidebarOnTeamLeave")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.ClearSidebarOnTeamLeave(userId, teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) CountPostsAfter(channelId string, timestamp int64, userId string) (int, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.CountPostsAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.CountPostsAfter(channelId, timestamp, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) CreateDirectChannel(userId *model.User, otherUserId *model.User) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.CreateDirectChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.CreateDirectChannel(userId, otherUserId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) CreateInitialSidebarCategories(userId string, teamId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.CreateInitialSidebarCategories")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.CreateInitialSidebarCategories(userId, teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) CreateSidebarCategory(userId string, teamId string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.CreateSidebarCategory")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.CreateSidebarCategory(userId, teamId, newCategory)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) Delete(channelId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.Delete(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) DeleteSidebarCategory(categoryId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.DeleteSidebarCategory")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.DeleteSidebarCategory(categoryId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) DeleteSidebarChannelsByPreferences(preferences *model.Preferences) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.DeleteSidebarChannelsByPreferences")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.DeleteSidebarChannelsByPreferences(preferences)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) Get(id string, allowFromCache bool) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.Get(id, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAll(teamId string) ([]*model.Channel, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAll(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllChannelMembersForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllChannelMembersForUser(userId, allowFromCache, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllChannelMembersNotifyPropsForChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllChannelMembersNotifyPropsForChannel(channelId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllChannels(page int, perPage int, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllChannels(page, perPage, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllChannelsCount(opts store.ChannelSearchOpts) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllChannelsCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllChannelsCount(opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) ([]*model.ChannelForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllChannelsForExportAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllChannelsForExportAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) ([]*model.DirectChannelForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetAllDirectChannelsForExportAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetAllDirectChannelsForExportAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetByName(team_id string, name string, allowFromCache bool) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetByName(team_id, name, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetByNameIncludeDeleted")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetByNameIncludeDeleted(team_id, name, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) ([]*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetByNames")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetByNames(team_id, names, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelCounts(teamId string, userId string) (*model.ChannelCounts, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelCounts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelCounts(teamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelMembersForExport(userId string, teamId string) ([]*model.ChannelMemberForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelMembersForExport")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelMembersForExport(userId, teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelMembersTimezones(channelId string) ([]model.StringMap, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelMembersTimezones")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelMembersTimezones(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelUnread(channelId string, userId string) (*model.ChannelUnread, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelUnread")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelUnread(channelId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannels(teamId string, userId string, includeDeleted bool, lastDeleteAt int) (*model.ChannelList, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannels(teamId, userId, includeDeleted, lastDeleteAt)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.Channel, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelsBatchForIndexing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelsBatchForIndexing(startTime, endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelsByIds(channelIds []string, includeDeleted bool) ([]*model.Channel, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelsByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelsByIds(channelIds, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) (model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetChannelsByScheme")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetChannelsByScheme(schemeId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetDeleted(team_id string, offset int, limit int, userId string) (*model.ChannelList, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetDeleted")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetDeleted(team_id, offset, limit, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetDeletedByName(team_id string, name string) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetDeletedByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetDeletedByName(team_id, name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetForPost(postId string) (*model.Channel, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetForPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetForPost(postId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetFromMaster(id string) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetFromMaster")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetFromMaster(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetGuestCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetGuestCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetGuestCount(channelId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMember(channelId string, userId string) (*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMember(channelId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMemberCount(channelId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMemberCountFromCache(channelId string) int64 {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMemberCountFromCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.ChannelStore.GetMemberCountFromCache(channelId)
	return result
}

func (s *OpenTracingLayerChannelStore) GetMemberCountsByGroup(channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMemberCountsByGroup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMemberCountsByGroup(channelID, includeTimezones)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMemberForPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMemberForPost(postId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMembers(channelId string, offset int, limit int) (*model.ChannelMembers, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMembers(channelId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMembersByIds(channelId string, userIds []string) (*model.ChannelMembers, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMembersByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMembersByIds(channelId, userIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMembersForUser(teamId string, userId string) (*model.ChannelMembers, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMembersForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMembersForUser(teamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMembersForUserWithPagination(teamId string, userId string, page int, perPage int) (*model.ChannelMembers, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMembersForUserWithPagination")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMembersForUserWithPagination(teamId, userId, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) (*model.ChannelList, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetMoreChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetMoreChannels(teamId, userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetPinnedPostCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetPinnedPostCount(channelId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetPinnedPosts(channelId string) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetPinnedPosts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetPinnedPosts(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetPrivateChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetPrivateChannelsForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetPrivateChannelsForTeam(teamId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetPublicChannelsByIdsForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetPublicChannelsByIdsForTeam(teamId, channelIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetPublicChannelsForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetPublicChannelsForTeam(teamId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetSidebarCategories(userId string, teamId string) (*model.OrderedSidebarCategories, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetSidebarCategories")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetSidebarCategories(userId, teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetSidebarCategory(categoryId string) (*model.SidebarCategoryWithChannels, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetSidebarCategory")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetSidebarCategory(categoryId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetSidebarCategoryOrder(userId string, teamId string) ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetSidebarCategoryOrder")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetSidebarCategoryOrder(userId, teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GetTeamChannels(teamId string) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GetTeamChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GetTeamChannels(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) GroupSyncedChannelCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.GroupSyncedChannelCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.GroupSyncedChannelCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) IncrementMentionCount(channelId string, userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.IncrementMentionCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.IncrementMentionCount(channelId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateAllChannelMembersForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateAllChannelMembersForUser(userId)

}

func (s *OpenTracingLayerChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateCacheForChannelMembersNotifyProps")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateCacheForChannelMembersNotifyProps(channelId)

}

func (s *OpenTracingLayerChannelStore) InvalidateChannel(id string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateChannel(id)

}

func (s *OpenTracingLayerChannelStore) InvalidateChannelByName(teamId string, name string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateChannelByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateChannelByName(teamId, name)

}

func (s *OpenTracingLayerChannelStore) InvalidateGuestCount(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateGuestCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateGuestCount(channelId)

}

func (s *OpenTracingLayerChannelStore) InvalidateMemberCount(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidateMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidateMemberCount(channelId)

}

func (s *OpenTracingLayerChannelStore) InvalidatePinnedPostCount(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.InvalidatePinnedPostCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.ChannelStore.InvalidatePinnedPostCount(channelId)

}

func (s *OpenTracingLayerChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.IsUserInChannelUseCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.ChannelStore.IsUserInChannelUseCache(userId, channelId)
	return result
}

func (s *OpenTracingLayerChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) (map[string]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.MigrateChannelMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.MigrateChannelMembers(fromChannelId, fromUserId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) MigratePublicChannels() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.MigratePublicChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.MigratePublicChannels()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) PermanentDelete(channelId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.PermanentDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.PermanentDelete(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) PermanentDeleteByTeam(teamId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.PermanentDeleteByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.PermanentDeleteByTeam(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) PermanentDeleteMembersByChannel(channelId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.PermanentDeleteMembersByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.PermanentDeleteMembersByChannel(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) PermanentDeleteMembersByUser(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.PermanentDeleteMembersByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.PermanentDeleteMembersByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) RemoveAllDeactivatedMembers(channelId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.RemoveAllDeactivatedMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.RemoveAllDeactivatedMembers(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) RemoveMember(channelId string, userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.RemoveMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.RemoveMember(channelId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) RemoveMembers(channelId string, userIds []string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.RemoveMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.RemoveMembers(channelId, userIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) ResetAllChannelSchemes() *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.ResetAllChannelSchemes")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.ResetAllChannelSchemes()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) Restore(channelId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.Restore")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.Restore(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.Save(channel, maxChannelsPerTeam)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SaveDirectChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SaveDirectChannel(channel, member1, member2)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SaveMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SaveMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SaveMember(member)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SaveMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SaveMultipleMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SaveMultipleMembers(members)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SearchAllChannels(term string, opts store.ChannelSearchOpts) (*model.ChannelListWithTeamData, int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchAllChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, resultVar1, err := s.ChannelStore.SearchAllChannels(term, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, resultVar1, err
}

func (s *OpenTracingLayerChannelStore) SearchArchivedInTeam(teamId string, term string, userId string) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchArchivedInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SearchArchivedInTeam(teamId, term, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SearchForUserInTeam(userId string, teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchForUserInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SearchForUserInTeam(userId, teamId, term, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SearchGroupChannels(userId string, term string) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchGroupChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SearchGroupChannels(userId, term)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SearchInTeam(teamId, term, includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SearchMore(userId string, teamId string, term string) (*model.ChannelList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SearchMore")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.SearchMore(userId, teamId, term)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.SetDeleteAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.SetDeleteAt(channelId, deleteAt, updateAt)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) Update(channel *model.Channel) (*model.Channel, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.Update(channel)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateLastViewedAt(channelIds []string, userId string) (map[string]int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateLastViewedAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UpdateLastViewedAt(channelIds, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateLastViewedAtPost(unreadPost *model.Post, userID string, mentionCount int) (*model.ChannelUnreadAt, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateLastViewedAtPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UpdateLastViewedAtPost(unreadPost, userID, mentionCount)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateMember(member *model.ChannelMember) (*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UpdateMember(member)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateMembersRole(channelID string, userIDs []string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateMembersRole")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.UpdateMembersRole(channelID, userIDs)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) UpdateMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateMultipleMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UpdateMultipleMembers(members)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateSidebarCategories(userId string, teamId string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateSidebarCategories")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UpdateSidebarCategories(userId, teamId, categories)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelStore) UpdateSidebarCategoryOrder(userId string, teamId string, categoryOrder []string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateSidebarCategoryOrder")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.UpdateSidebarCategoryOrder(userId, teamId, categoryOrder)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) UpdateSidebarChannelCategoryOnMove(channel *model.Channel, newTeamId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateSidebarChannelCategoryOnMove")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.UpdateSidebarChannelCategoryOnMove(channel, newTeamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) UpdateSidebarChannelsByPreferences(preferences *model.Preferences) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UpdateSidebarChannelsByPreferences")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelStore.UpdateSidebarChannelsByPreferences(preferences)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelStore) UserBelongsToChannels(userId string, channelIds []string) (bool, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelStore.UserBelongsToChannels")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelStore.UserBelongsToChannels(userId, channelIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) ([]*model.ChannelMemberHistoryResult, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelMemberHistoryStore.GetUsersInChannelDuring")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelMemberHistoryStore.GetUsersInChannelDuring(startTime, endTime, channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelMemberHistoryStore.LogJoinEvent")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelMemberHistoryStore.LogJoinEvent(userId, channelId, joinTime)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelMemberHistoryStore.LogLeaveEvent")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ChannelMemberHistoryStore.LogLeaveEvent(userId, channelId, leaveTime)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ChannelMemberHistoryStore.PermanentDeleteBatch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ChannelMemberHistoryStore.PermanentDeleteBatch(endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerClusterDiscoveryStore) Cleanup() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.Cleanup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ClusterDiscoveryStore.Cleanup()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ClusterDiscoveryStore.Delete(discovery)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.Exists")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ClusterDiscoveryStore.Exists(discovery)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ClusterDiscoveryStore.GetAll(discoveryType, clusterName)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ClusterDiscoveryStore.Save(discovery)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ClusterDiscoveryStore.SetLastPingAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ClusterDiscoveryStore.SetLastPingAt(discovery)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerCommandStore) AnalyticsCommandCount(teamId string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.AnalyticsCommandCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.AnalyticsCommandCount(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandStore) Delete(commandId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.CommandStore.Delete(commandId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerCommandStore) Get(id string) (*model.Command, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandStore) GetByTeam(teamId string) ([]*model.Command, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.GetByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.GetByTeam(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandStore) GetByTrigger(teamId string, trigger string) (*model.Command, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.GetByTrigger")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.GetByTrigger(teamId, trigger)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandStore) PermanentDeleteByTeam(teamId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.PermanentDeleteByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.CommandStore.PermanentDeleteByTeam(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerCommandStore) PermanentDeleteByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.PermanentDeleteByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.CommandStore.PermanentDeleteByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerCommandStore) Save(webhook *model.Command) (*model.Command, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.Save(webhook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandStore) Update(hook *model.Command) (*model.Command, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandStore.Update(hook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandWebhookStore) Cleanup() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandWebhookStore.Cleanup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.CommandWebhookStore.Cleanup()

}

func (s *OpenTracingLayerCommandWebhookStore) Get(id string) (*model.CommandWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandWebhookStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandWebhookStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandWebhookStore) Save(webhook *model.CommandWebhook) (*model.CommandWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandWebhookStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.CommandWebhookStore.Save(webhook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerCommandWebhookStore) TryUse(id string, limit int) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "CommandWebhookStore.TryUse")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.CommandWebhookStore.TryUse(id, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.ComplianceExport")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.ComplianceExport(compliance)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerComplianceStore) Get(id string) (*model.Compliance, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerComplianceStore) GetAll(offset int, limit int) (model.Compliances, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.GetAll(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.MessageExport")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.MessageExport(after, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.Save(compliance)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ComplianceStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ComplianceStore.Update(compliance)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) Delete(emoji *model.Emoji, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.EmojiStore.Delete(emoji, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerEmojiStore) Get(id string, allowFromCache bool) (*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.Get(id, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) GetByName(name string, allowFromCache bool) (*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.GetByName(name, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) GetList(offset int, limit int, sort string) ([]*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.GetList")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.GetList(offset, limit, sort)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) GetMultipleByName(names []string) ([]*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.GetMultipleByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.GetMultipleByName(names)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.Save(emoji)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerEmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "EmojiStore.Search")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.EmojiStore.Search(name, prefixOnly, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) AttachToPost(fileId string, postId string, creatorId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.AttachToPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.FileInfoStore.AttachToPost(fileId, postId, creatorId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerFileInfoStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.FileInfoStore.ClearCaches()

}

func (s *OpenTracingLayerFileInfoStore) DeleteForPost(postId string) (string, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.DeleteForPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.DeleteForPost(postId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) Get(id string) (*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) GetByPath(path string) (*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.GetByPath")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.GetByPath(path)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) GetForPost(postId string, readFromMaster bool, includeDeleted bool, allowFromCache bool) ([]*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.GetForPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.GetForPost(postId, readFromMaster, includeDeleted, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) GetForUser(userId string) ([]*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.GetForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.GetForUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) GetWithOptions(page int, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.GetWithOptions")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.GetWithOptions(page, perPage, opt)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) InvalidateFileInfosForPostCache(postId string, deleted bool) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.InvalidateFileInfosForPostCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.FileInfoStore.InvalidateFileInfosForPostCache(postId, deleted)

}

func (s *OpenTracingLayerFileInfoStore) PermanentDelete(fileId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.PermanentDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.FileInfoStore.PermanentDelete(fileId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerFileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.PermanentDeleteBatch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.PermanentDeleteBatch(endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) PermanentDeleteByUser(userId string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.PermanentDeleteByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.PermanentDeleteByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerFileInfoStore) Save(info *model.FileInfo) (*model.FileInfo, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "FileInfoStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.FileInfoStore.Save(info)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) AdminRoleGroupsForSyncableMember(userID string, syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.AdminRoleGroupsForSyncableMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.AdminRoleGroupsForSyncableMember(userID, syncableID, syncableType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.ChannelMembersMinusGroupMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.ChannelMembersMinusGroupMembers(channelID, groupIDs, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) ChannelMembersToAdd(since int64, channelID *string) ([]*model.UserChannelIDPair, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.ChannelMembersToAdd")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.ChannelMembersToAdd(since, channelID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.ChannelMembersToRemove")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.ChannelMembersToRemove(channelID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.CountChannelMembersMinusGroupMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.CountChannelMembersMinusGroupMembers(channelID, groupIDs)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.CountGroupsByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.CountGroupsByChannel(channelId, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.CountGroupsByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.CountGroupsByTeam(teamId, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.CountTeamMembersMinusGroupMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.CountTeamMembersMinusGroupMembers(teamID, groupIDs)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) Create(group *model.Group) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.Create")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.Create(group)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.CreateGroupSyncable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.CreateGroupSyncable(groupSyncable)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) Delete(groupID string) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.Delete(groupID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.DeleteGroupSyncable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.DeleteGroupSyncable(groupID, syncableID, syncableType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.DeleteMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.DeleteMember(groupID, userID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) DistinctGroupMemberCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.DistinctGroupMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.DistinctGroupMemberCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) Get(groupID string) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.Get(groupID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetAllBySource")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetAllBySource(groupSource)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetAllGroupSyncablesByGroupId")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetByIDs(groupIDs []string) ([]*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetByIDs")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetByIDs(groupIDs)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetByName(name, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetByRemoteID")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetByRemoteID(remoteID, groupSource)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetByUser(userId string) ([]*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetGroupSyncable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetGroupSyncable(groupID, syncableID, syncableType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetGroups(page int, perPage int, opts model.GroupSearchOpts) ([]*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetGroups")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetGroups(page, perPage, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetGroupsAssociatedToChannelsByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetGroupsAssociatedToChannelsByTeam(teamId, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetGroupsByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetGroupsByChannel(channelId, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetGroupsByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetGroupsByTeam(teamId, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetMemberCount(groupID string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetMemberCount(groupID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetMemberUsers(groupID string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetMemberUsers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetMemberUsers(groupID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetMemberUsersInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetMemberUsersInTeam(groupID, teamID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetMemberUsersNotInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetMemberUsersNotInChannel(groupID, channelID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GetMemberUsersPage")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GetMemberUsersPage(groupID, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GroupChannelCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GroupChannelCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GroupChannelCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GroupCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GroupCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GroupCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GroupCountWithAllowReference() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GroupCountWithAllowReference")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GroupCountWithAllowReference()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GroupMemberCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GroupMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GroupMemberCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) GroupTeamCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.GroupTeamCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.GroupTeamCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) PermanentDeleteMembersByUser(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.PermanentDeleteMembersByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.GroupStore.PermanentDeleteMembersByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerGroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.PermittedSyncableAdmins")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.PermittedSyncableAdmins(syncableID, syncableType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.TeamMembersMinusGroupMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.TeamMembersMinusGroupMembers(teamID, groupIDs, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) TeamMembersToAdd(since int64, teamID *string) ([]*model.UserTeamIDPair, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.TeamMembersToAdd")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.TeamMembersToAdd(since, teamID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.TeamMembersToRemove")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.TeamMembersToRemove(teamID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) Update(group *model.Group) (*model.Group, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.Update(group)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.UpdateGroupSyncable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.UpdateGroupSyncable(groupSyncable)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerGroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "GroupStore.UpsertMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.GroupStore.UpsertMember(groupID, userID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) Delete(id string) (string, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.Delete(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) Get(id string) (*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetAllByStatus(status string) ([]*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetAllByStatus")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetAllByStatus(status)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetAllByType(jobType string) ([]*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetAllByType")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetAllByType(jobType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetAllByTypePage(jobType string, offset int, limit int) ([]*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetAllByTypePage")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetAllByTypePage(jobType, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetAllPage(offset int, limit int) ([]*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetAllPage")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetAllPage(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetCountByStatusAndType(status string, jobType string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetCountByStatusAndType")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetCountByStatusAndType(status, jobType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetNewestJobByStatusAndType(status string, jobType string) (*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetNewestJobByStatusAndType")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetNewestJobByStatusAndType(status, jobType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) GetNewestJobByStatusesAndType(statuses []string, jobType string) (*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.GetNewestJobByStatusesAndType")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.GetNewestJobByStatusesAndType(statuses, jobType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) Save(job *model.Job) (*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.Save(job)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) UpdateOptimistically(job *model.Job, currentStatus string) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.UpdateOptimistically")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.UpdateOptimistically(job, currentStatus)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) UpdateStatus(id string, status string) (*model.Job, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.UpdateStatus")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.UpdateStatus(id, status)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerJobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "JobStore.UpdateStatusOptimistically")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.JobStore.UpdateStatusOptimistically(id, currentStatus, newStatus)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerLicenseStore) Get(id string) (*model.LicenseRecord, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "LicenseStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.LicenseStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "LicenseStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.LicenseStore.Save(license)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerLinkMetadataStore) Get(url string, timestamp int64) (*model.LinkMetadata, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "LinkMetadataStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.LinkMetadataStore.Get(url, timestamp)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerLinkMetadataStore) Save(linkMetadata *model.LinkMetadata) (*model.LinkMetadata, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "LinkMetadataStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.LinkMetadataStore.Save(linkMetadata)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) DeleteApp(id string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.DeleteApp")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.OAuthStore.DeleteApp(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerOAuthStore) GetAccessData(token string) (*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAccessData(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetAccessDataByRefreshToken(token string) (*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAccessDataByRefreshToken")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAccessDataByRefreshToken(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetAccessDataByUserForApp(userId string, clientId string) ([]*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAccessDataByUserForApp")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAccessDataByUserForApp(userId, clientId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetApp(id string) (*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetApp")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetApp(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetAppByUser(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAppByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAppByUser(userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetApps(offset int, limit int) ([]*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetApps")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetApps(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetAuthData(code string) (*model.AuthData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAuthData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAuthData(code)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetAuthorizedApps(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetAuthorizedApps")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetAuthorizedApps(userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) GetPreviousAccessData(userId string, clientId string) (*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.GetPreviousAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.GetPreviousAccessData(userId, clientId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) PermanentDeleteAuthDataByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.PermanentDeleteAuthDataByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.OAuthStore.PermanentDeleteAuthDataByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerOAuthStore) RemoveAccessData(token string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.RemoveAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.OAuthStore.RemoveAccessData(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerOAuthStore) RemoveAllAccessData() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.RemoveAllAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.OAuthStore.RemoveAllAccessData()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerOAuthStore) RemoveAuthData(code string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.RemoveAuthData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.OAuthStore.RemoveAuthData(code)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerOAuthStore) SaveAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.SaveAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.SaveAccessData(accessData)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) SaveApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.SaveApp")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.SaveApp(app)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) SaveAuthData(authData *model.AuthData) (*model.AuthData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.SaveAuthData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.SaveAuthData(authData)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) UpdateAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.UpdateAccessData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.UpdateAccessData(accessData)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerOAuthStore) UpdateApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "OAuthStore.UpdateApp")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.OAuthStore.UpdateApp(app)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) CompareAndDelete(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.CompareAndDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.CompareAndDelete(keyVal, oldValue)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.CompareAndSet")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.CompareAndSet(keyVal, oldValue)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) Delete(pluginId string, key string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PluginStore.Delete(pluginId, key)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPluginStore) DeleteAllExpired() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.DeleteAllExpired")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PluginStore.DeleteAllExpired()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPluginStore) DeleteAllForPlugin(PluginId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.DeleteAllForPlugin")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PluginStore.DeleteAllForPlugin(PluginId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPluginStore) Get(pluginId string, key string) (*model.PluginKeyValue, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.Get(pluginId, key)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) List(pluginId string, page int, perPage int) ([]string, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.List")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.List(pluginId, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) (*model.PluginKeyValue, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.SaveOrUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.SaveOrUpdate(keyVal)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPluginStore) SetWithOptions(pluginId string, key string, value []byte, options model.PluginKVSetOptions) (bool, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PluginStore.SetWithOptions")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PluginStore.SetWithOptions(pluginId, key, value, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.AnalyticsPostCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.AnalyticsPostCount(teamId, mustHaveFile, mustHaveHashtag)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.AnalyticsPostCountsByDay")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.AnalyticsPostCountsByDay(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) (model.AnalyticsRows, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.AnalyticsUserCountsWithPostsByDay")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.AnalyticsUserCountsWithPostsByDay(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.PostStore.ClearCaches()

}

func (s *OpenTracingLayerPostStore) Delete(postId string, time int64, deleteByID string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PostStore.Delete(postId, time, deleteByID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPostStore) Get(id string, skipFetchThreads bool) (*model.PostList, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.Get(id, skipFetchThreads)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) ([]*model.DirectPostForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetDirectPostParentsForExportAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetDirectPostParentsForExportAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetEtag(channelId string, allowFromCache bool) string {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetEtag")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.PostStore.GetEtag(channelId, allowFromCache)
	return result
}

func (s *OpenTracingLayerPostStore) GetFlaggedPosts(userId string, offset int, limit int) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetFlaggedPosts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetFlaggedPosts(userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetFlaggedPostsForChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetFlaggedPostsForChannel(userId, channelId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetFlaggedPostsForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	span.SetTag("userId", userId)

	span.SetTag("teamId", teamId)

	span.SetTag("offset", offset)

	span.SetTag("limit", limit)

	defer span.Finish()
	result, err := s.PostStore.GetFlaggedPostsForTeam(userId, teamId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetMaxPostSize() int {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetMaxPostSize")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.PostStore.GetMaxPostSize()
	return result
}

func (s *OpenTracingLayerPostStore) GetOldest() (*model.Post, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetOldest")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetOldest()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetOldestEntityCreationTime() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetOldestEntityCreationTime")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetOldestEntityCreationTime()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetParentsForExportAfter(limit int, afterId string) ([]*model.PostForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetParentsForExportAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetParentsForExportAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostAfterTime(channelId string, time int64) (*model.Post, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostAfterTime")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostAfterTime(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostIdAfterTime(channelId string, time int64) (string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostIdAfterTime")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostIdAfterTime(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostIdBeforeTime(channelId string, time int64) (string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostIdBeforeTime")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostIdBeforeTime(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPosts(options model.GetPostsOptions, allowFromCache bool) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPosts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPosts(options, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsAfter(options model.GetPostsOptions) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsAfter(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.PostForIndexing, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsBatchForIndexing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsBatchForIndexing(startTime, endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsBefore(options model.GetPostsOptions) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsBefore")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsBefore(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsByIds(postIds []string) ([]*model.Post, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsByIds(postIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsCreatedAt(channelId string, time int64) ([]*model.Post, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsCreatedAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsCreatedAt(channelId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetPostsSince")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetPostsSince(options, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetRepliesForExport(parentId string) ([]*model.ReplyForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetRepliesForExport")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetRepliesForExport(parentId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) GetSingle(id string) (*model.Post, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.GetSingle")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.GetSingle(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) InvalidateLastPostTimeCache(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.InvalidateLastPostTimeCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.PostStore.InvalidateLastPostTimeCache(channelId)

}

func (s *OpenTracingLayerPostStore) Overwrite(post *model.Post) (*model.Post, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Overwrite")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.Overwrite(post)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.OverwriteMultiple")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, resultVar1, err := s.PostStore.OverwriteMultiple(posts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, resultVar1, err
}

func (s *OpenTracingLayerPostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.PermanentDeleteBatch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.PermanentDeleteBatch(endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) PermanentDeleteByChannel(channelId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.PermanentDeleteByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PostStore.PermanentDeleteByChannel(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPostStore) PermanentDeleteByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.PermanentDeleteByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PostStore.PermanentDeleteByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPostStore) Save(post *model.Post) (*model.Post, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.Save(post)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.SaveMultiple")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, resultVar1, err := s.PostStore.SaveMultiple(posts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, resultVar1, err
}

func (s *OpenTracingLayerPostStore) Search(teamId string, userId string, params *model.SearchParams) (*model.PostList, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Search")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.Search(teamId, userId, params)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) SearchPostsInTeamForUser(paramsList []*model.SearchParams, userId string, teamId string, isOrSearch bool, includeDeletedChannels bool, page int, perPage int) (*model.PostSearchResults, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.SearchPostsInTeamForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.SearchPostsInTeamForUser(paramsList, userId, teamId, isOrSearch, includeDeletedChannels, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPostStore) Update(newPost *model.Post, oldPost *model.Post) (*model.Post, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PostStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PostStore.Update(newPost, oldPost)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.CleanupFlagsBatch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PreferenceStore.CleanupFlagsBatch(limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPreferenceStore) Delete(userId string, category string, name string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PreferenceStore.Delete(userId, category, name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPreferenceStore) DeleteCategory(userId string, category string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.DeleteCategory")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PreferenceStore.DeleteCategory(userId, category)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPreferenceStore) DeleteCategoryAndName(category string, name string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.DeleteCategoryAndName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PreferenceStore.DeleteCategoryAndName(category, name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPreferenceStore) Get(userId string, category string, name string) (*model.Preference, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PreferenceStore.Get(userId, category, name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPreferenceStore) GetAll(userId string) (model.Preferences, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PreferenceStore.GetAll(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPreferenceStore) GetCategory(userId string, category string) (model.Preferences, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.GetCategory")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.PreferenceStore.GetCategory(userId, category)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerPreferenceStore) PermanentDeleteByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.PermanentDeleteByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PreferenceStore.PermanentDeleteByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerPreferenceStore) Save(preferences *model.Preferences) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "PreferenceStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.PreferenceStore.Save(preferences)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.BulkGetForPosts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ReactionStore.BulkGetForPosts(postIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ReactionStore.Delete(reaction)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerReactionStore) DeleteAllWithEmojiName(emojiName string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.DeleteAllWithEmojiName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.ReactionStore.DeleteAllWithEmojiName(emojiName)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.GetForPost")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ReactionStore.GetForPost(postId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.PermanentDeleteBatch")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ReactionStore.PermanentDeleteBatch(endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerReactionStore) Save(reaction *model.Reaction) (*model.Reaction, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "ReactionStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.ReactionStore.Save(reaction)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) AllChannelSchemeRoles() ([]*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.AllChannelSchemeRoles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.AllChannelSchemeRoles()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) ChannelHigherScopedPermissions(roleNames []string) (map[string]*model.RolePermissions, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.ChannelHigherScopedPermissions")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.ChannelHigherScopedPermissions(roleNames)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) ChannelRolesUnderTeamRole(roleName string) ([]*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.ChannelRolesUnderTeamRole")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.ChannelRolesUnderTeamRole(roleName)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) Delete(roleId string) (*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.Delete(roleId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) Get(roleId string) (*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.Get(roleId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) GetAll() ([]*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.GetAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) GetByName(name string) (*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.GetByName(name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) GetByNames(names []string) ([]*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.GetByNames")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.GetByNames(names)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerRoleStore) PermanentDeleteAll() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.PermanentDeleteAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.RoleStore.PermanentDeleteAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerRoleStore) Save(role *model.Role) (*model.Role, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "RoleStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.RoleStore.Save(role)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) CountByScope(scope string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.CountByScope")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.CountByScope(scope)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) CountWithoutPermission(scope string, permissionID string, roleScope model.RoleScope, roleType model.RoleType) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.CountWithoutPermission")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.CountWithoutPermission(scope, permissionID, roleScope, roleType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) Delete(schemeId string) (*model.Scheme, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.Delete(schemeId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) Get(schemeId string) (*model.Scheme, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.Get(schemeId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) GetAllPage(scope string, offset int, limit int) ([]*model.Scheme, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.GetAllPage")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.GetAllPage(scope, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) GetByName(schemeName string) (*model.Scheme, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.GetByName(schemeName)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSchemeStore) PermanentDeleteAll() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.PermanentDeleteAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SchemeStore.PermanentDeleteAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSchemeStore) Save(scheme *model.Scheme) (*model.Scheme, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SchemeStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SchemeStore.Save(scheme)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) AnalyticsSessionCount() (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.AnalyticsSessionCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.AnalyticsSessionCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) Cleanup(expiryTime int64, batchSize int64) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.Cleanup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.SessionStore.Cleanup(expiryTime, batchSize)

}

func (s *OpenTracingLayerSessionStore) Get(sessionIdOrToken string) (*model.Session, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.Get(sessionIdOrToken)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) GetSessions(userId string) ([]*model.Session, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.GetSessions")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.GetSessions(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) GetSessionsExpired(thresholdMillis int64, mobileOnly bool, unnotifiedOnly bool) ([]*model.Session, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.GetSessionsExpired")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.GetSessionsExpired(thresholdMillis, mobileOnly, unnotifiedOnly)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) GetSessionsWithActiveDeviceIds(userId string) ([]*model.Session, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.GetSessionsWithActiveDeviceIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.GetSessionsWithActiveDeviceIds(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) PermanentDeleteSessionsByUser(teamId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.PermanentDeleteSessionsByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.PermanentDeleteSessionsByUser(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) Remove(sessionIdOrToken string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.Remove")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.Remove(sessionIdOrToken)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) RemoveAllSessions() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.RemoveAllSessions")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.RemoveAllSessions()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) Save(session *model.Session) (*model.Session, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.Save(session)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) (string, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateDeviceId")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.UpdateDeviceId(id, deviceId, expiresAt)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSessionStore) UpdateExpiredNotify(sessionid string, notified bool) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateExpiredNotify")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.UpdateExpiredNotify(sessionid, notified)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) UpdateExpiresAt(sessionId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateExpiresAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.UpdateExpiresAt(sessionId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) UpdateLastActivityAt(sessionId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateLastActivityAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.UpdateLastActivityAt(sessionId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) UpdateProps(session *model.Session) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateProps")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SessionStore.UpdateProps(session)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSessionStore) UpdateRoles(userId string, roles string) (string, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SessionStore.UpdateRoles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SessionStore.UpdateRoles(userId, roles)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerStatusStore) Get(userId string) (*model.Status, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.StatusStore.Get(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerStatusStore) GetByIds(userIds []string) ([]*model.Status, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.GetByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.StatusStore.GetByIds(userIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerStatusStore) GetTotalActiveUsersCount() (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.GetTotalActiveUsersCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.StatusStore.GetTotalActiveUsersCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerStatusStore) ResetAll() error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.ResetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.StatusStore.ResetAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerStatusStore) SaveOrUpdate(status *model.Status) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.SaveOrUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.StatusStore.SaveOrUpdate(status)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "StatusStore.UpdateLastActivityAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.StatusStore.UpdateLastActivityAt(userId, lastActivityAt)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSystemStore) Get() (model.StringMap, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SystemStore.Get()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSystemStore) GetByName(name string) (*model.System, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SystemStore.GetByName(name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSystemStore) InsertIfExists(system *model.System) (*model.System, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.InsertIfExists")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SystemStore.InsertIfExists(system)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSystemStore) PermanentDeleteByName(name string) (*model.System, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.PermanentDeleteByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.SystemStore.PermanentDeleteByName(name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerSystemStore) Save(system *model.System) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SystemStore.Save(system)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSystemStore) SaveOrUpdate(system *model.System) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.SaveOrUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SystemStore.SaveOrUpdate(system)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerSystemStore) Update(system *model.System) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "SystemStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.SystemStore.Update(system)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.AnalyticsGetTeamCountForScheme")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.AnalyticsGetTeamCountForScheme(schemeId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) AnalyticsPrivateTeamCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.AnalyticsPrivateTeamCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.AnalyticsPrivateTeamCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) AnalyticsPublicTeamCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.AnalyticsPublicTeamCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.AnalyticsPublicTeamCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) AnalyticsTeamCount(includeDeleted bool) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.AnalyticsTeamCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.AnalyticsTeamCount(includeDeleted)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) ClearAllCustomRoleAssignments() *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.ClearAllCustomRoleAssignments")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.ClearAllCustomRoleAssignments()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.TeamStore.ClearCaches()

}

func (s *OpenTracingLayerTeamStore) Get(id string) (*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetActiveMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetActiveMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetActiveMemberCount(teamId, restrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAll() ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllForExportAfter(limit int, afterId string) ([]*model.TeamForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllForExportAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllForExportAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllPage(offset int, limit int) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllPage")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllPage(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllPrivateTeamListing() ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllPrivateTeamListing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllPrivateTeamListing()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllPrivateTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllPrivateTeamPageListing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllPrivateTeamPageListing(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllPublicTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllPublicTeamPageListing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllPublicTeamPageListing(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllTeamListing() ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllTeamListing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllTeamListing()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetAllTeamPageListing(offset int, limit int) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetAllTeamPageListing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetAllTeamPageListing(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetByInviteId(inviteId string) (*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetByInviteId")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetByInviteId(inviteId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetByName(name string) (*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetByName")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetByName(name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetByNames(name []string) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetByNames")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetByNames(name)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetChannelUnreadsForAllTeams")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetChannelUnreadsForAllTeams(excludeTeamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetChannelUnreadsForTeam(teamId string, userId string) ([]*model.ChannelUnread, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetChannelUnreadsForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetChannelUnreadsForTeam(teamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetMember(teamId string, userId string) (*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetMember(teamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetMembers(teamId string, offset int, limit int, teamMembersGetOptions *model.TeamMembersGetOptions) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetMembers(teamId, offset, limit, teamMembersGetOptions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetMembersByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetMembersByIds(teamId, userIds, restrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTeamMembersForExport(userId string) ([]*model.TeamMemberForExport, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTeamMembersForExport")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTeamMembersForExport(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTeamsByScheme")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTeamsByScheme(schemeId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTeamsByUserId(userId string) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTeamsByUserId")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTeamsByUserId(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTeamsForUser(userId string) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTeamsForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTeamsForUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTeamsForUserWithPagination")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTeamsForUserWithPagination(userId, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetTotalMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetTotalMemberCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetTotalMemberCount(teamId, restrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GetUserTeamIds(userId string, allowFromCache bool) ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GetUserTeamIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GetUserTeamIds(userId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) GroupSyncedTeamCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.GroupSyncedTeamCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.GroupSyncedTeamCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) InvalidateAllTeamIdsForUser(userId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.InvalidateAllTeamIdsForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.TeamStore.InvalidateAllTeamIdsForUser(userId)

}

func (s *OpenTracingLayerTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) (map[string]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.MigrateTeamMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.MigrateTeamMembers(fromTeamId, fromUserId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) PermanentDelete(teamId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.PermanentDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.PermanentDelete(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) RemoveAllMembersByTeam(teamId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.RemoveAllMembersByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.RemoveAllMembersByTeam(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) RemoveAllMembersByUser(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.RemoveAllMembersByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.RemoveAllMembersByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) RemoveMember(teamId string, userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.RemoveMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.RemoveMember(teamId, userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) RemoveMembers(teamId string, userIds []string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.RemoveMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.RemoveMembers(teamId, userIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) ResetAllTeamSchemes() *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.ResetAllTeamSchemes")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.ResetAllTeamSchemes()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) Save(team *model.Team) (*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.Save(team)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) (*model.TeamMember, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SaveMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.SaveMember(member, maxUsersPerTeam)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) SaveMultipleMembers(members []*model.TeamMember, maxUsersPerTeam int) ([]*model.TeamMember, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SaveMultipleMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.SaveMultipleMembers(members, maxUsersPerTeam)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) SearchAll(term string, opts *model.TeamSearch) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SearchAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.SearchAll(term, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) SearchAllPaged(term string, opts *model.TeamSearch) ([]*model.Team, int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SearchAllPaged")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, resultVar1, err := s.TeamStore.SearchAllPaged(term, opts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, resultVar1, err
}

func (s *OpenTracingLayerTeamStore) SearchOpen(term string) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SearchOpen")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.SearchOpen(term)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) SearchPrivate(term string) ([]*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.SearchPrivate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.SearchPrivate(term)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.Update(team)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.UpdateLastTeamIconUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.UpdateLastTeamIconUpdate(teamId, curTime)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) UpdateMember(member *model.TeamMember) (*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.UpdateMember")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.UpdateMember(member)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) UpdateMembersRole(teamID string, userIDs []string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.UpdateMembersRole")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TeamStore.UpdateMembersRole(teamID, userIDs)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTeamStore) UpdateMultipleMembers(members []*model.TeamMember) ([]*model.TeamMember, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.UpdateMultipleMembers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.UpdateMultipleMembers(members)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTeamStore) UserBelongsToTeams(userId string, teamIds []string) (bool, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TeamStore.UserBelongsToTeams")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TeamStore.UserBelongsToTeams(userId, teamIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTermsOfServiceStore) Get(id string, allowFromCache bool) (*model.TermsOfService, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TermsOfServiceStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TermsOfServiceStore.Get(id, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTermsOfServiceStore) GetLatest(allowFromCache bool) (*model.TermsOfService, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TermsOfServiceStore.GetLatest")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TermsOfServiceStore.GetLatest(allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTermsOfServiceStore) Save(termsOfService *model.TermsOfService) (*model.TermsOfService, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TermsOfServiceStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TermsOfServiceStore.Save(termsOfService)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTokenStore) Cleanup() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TokenStore.Cleanup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.TokenStore.Cleanup()

}

func (s *OpenTracingLayerTokenStore) Delete(token string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TokenStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TokenStore.Delete(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTokenStore) GetByToken(token string) (*model.Token, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TokenStore.GetByToken")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.TokenStore.GetByToken(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerTokenStore) RemoveAllTokensByType(tokenType string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TokenStore.RemoveAllTokensByType")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TokenStore.RemoveAllTokensByType(tokenType)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerTokenStore) Save(recovery *model.Token) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "TokenStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.TokenStore.Save(recovery)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) AnalyticsActiveCount(time int64, options model.UserCountOptions) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.AnalyticsActiveCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.AnalyticsActiveCount(time, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) AnalyticsGetGuestCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.AnalyticsGetGuestCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.AnalyticsGetGuestCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) AnalyticsGetInactiveUsersCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.AnalyticsGetInactiveUsersCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.AnalyticsGetInactiveUsersCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) AnalyticsGetSystemAdminCount() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.AnalyticsGetSystemAdminCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.AnalyticsGetSystemAdminCount()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) AutocompleteUsersInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) (*model.UserAutocompleteInChannel, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.AutocompleteUsersInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.AutocompleteUsersInChannel(teamId, channelId, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) ClearAllCustomRoleAssignments() *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.ClearAllCustomRoleAssignments")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.ClearAllCustomRoleAssignments()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.UserStore.ClearCaches()

}

func (s *OpenTracingLayerUserStore) Count(options model.UserCountOptions) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.Count")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.Count(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) DeactivateGuests() ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.DeactivateGuests")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.DeactivateGuests()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) DemoteUserToGuest(userID string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.DemoteUserToGuest")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.DemoteUserToGuest(userID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) Get(id string) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.Get(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAll() ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAll()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAllAfter(limit int, afterId string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAllAfter")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAllAfter(limit, afterId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAllNotInAuthService(authServices []string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAllNotInAuthService")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAllNotInAuthService(authServices)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAllProfiles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAllProfiles(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) (map[string]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAllProfilesInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAllProfilesInChannel(channelId, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAllUsingAuthService(authService string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAllUsingAuthService")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAllUsingAuthService(authService)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetAnyUnreadPostCountForChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetAnyUnreadPostCountForChannel(userId, channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetByAuth(authData *string, authService string) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetByAuth")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetByAuth(authData, authService)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetByEmail(email string) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetByEmail")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetByEmail(email)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetByUsername(username string) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetByUsername")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetByUsername(username)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetChannelGroupUsers(channelID string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetChannelGroupUsers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetChannelGroupUsers(channelID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetEtagForAllProfiles() string {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetEtagForAllProfiles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.UserStore.GetEtagForAllProfiles()
	return result
}

func (s *OpenTracingLayerUserStore) GetEtagForProfiles(teamId string) string {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetEtagForProfiles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.UserStore.GetEtagForProfiles(teamId)
	return result
}

func (s *OpenTracingLayerUserStore) GetEtagForProfilesNotInTeam(teamId string) string {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetEtagForProfilesNotInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result := s.UserStore.GetEtagForProfilesNotInTeam(teamId)
	return result
}

func (s *OpenTracingLayerUserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetForLogin")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetForLogin(loginId, allowSignInWithUsername, allowSignInWithEmail)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetKnownUsers(userID string) ([]string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetKnownUsers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetKnownUsers(userID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetNewUsersForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetNewUsersForTeam(teamId, offset, limit, viewRestrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfileByGroupChannelIdsForUser(userId string, channelIds []string) (map[string][]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfileByGroupChannelIdsForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfileByGroupChannelIdsForUser(userId, channelIds)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfileByIds(userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfileByIds")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfileByIds(userIds, options, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfiles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfiles(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesByUsernames")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesByUsernames(usernames, viewRestrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesInChannel(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesInChannel(channelId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesInChannelByStatus")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesInChannelByStatus(channelId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesNotInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesNotInChannel(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesNotInTeam(teamId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesNotInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesNotInTeam(teamId, groupConstrained, offset, limit, viewRestrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetProfilesWithoutTeam(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetProfilesWithoutTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetProfilesWithoutTeam(options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetRecentlyActiveUsersForTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetRecentlyActiveUsersForTeam(teamId, offset, limit, viewRestrictions)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetSystemAdminProfiles() (map[string]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetSystemAdminProfiles")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetSystemAdminProfiles()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetTeamGroupUsers(teamID string) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetTeamGroupUsers")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetTeamGroupUsers(teamID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetUnreadCount(userId string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetUnreadCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetUnreadCount(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetUnreadCountForChannel(userId string, channelId string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetUnreadCountForChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetUnreadCountForChannel(userId, channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.UserForIndexing, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.GetUsersBatchForIndexing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.GetUsersBatchForIndexing(startTime, endTime, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) InferSystemInstallDate() (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.InferSystemInstallDate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.InferSystemInstallDate()
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) InvalidateProfileCacheForUser(userId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.InvalidateProfileCacheForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.UserStore.InvalidateProfileCacheForUser(userId)

}

func (s *OpenTracingLayerUserStore) InvalidateProfilesInChannelCache(channelId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.InvalidateProfilesInChannelCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.UserStore.InvalidateProfilesInChannelCache(channelId)

}

func (s *OpenTracingLayerUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.InvalidateProfilesInChannelCacheByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.UserStore.InvalidateProfilesInChannelCacheByUser(userId)

}

func (s *OpenTracingLayerUserStore) PermanentDelete(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.PermanentDelete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.PermanentDelete(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) PromoteGuestToUser(userID string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.PromoteGuestToUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.PromoteGuestToUser(userID)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) ResetLastPictureUpdate(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.ResetLastPictureUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.ResetLastPictureUpdate(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) Save(user *model.User) (*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.Save(user)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) Search(teamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.Search")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.Search(teamId, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.SearchInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.SearchInChannel(channelId, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) SearchInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.SearchInGroup")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.SearchInGroup(groupID, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.SearchNotInChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.SearchNotInChannel(teamId, channelId, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.SearchNotInTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.SearchNotInTeam(notInTeamId, term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.SearchWithoutTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.SearchWithoutTeam(term, options)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) Update(user *model.User, allowRoleUpdate bool) (*model.UserUpdate, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.Update")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.Update(user, allowRoleUpdate)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) (string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateAuthData")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.UpdateAuthData(userId, service, authData, email, resetMfa)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateFailedPasswordAttempts")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.UpdateFailedPasswordAttempts(userId, attempts)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) UpdateLastPictureUpdate(userId string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateLastPictureUpdate")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.UpdateLastPictureUpdate(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) UpdateMfaActive(userId string, active bool) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateMfaActive")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.UpdateMfaActive(userId, active)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) UpdateMfaSecret(userId string, secret string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateMfaSecret")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.UpdateMfaSecret(userId, secret)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) UpdatePassword(userId string, newPassword string) *model.AppError {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdatePassword")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserStore.UpdatePassword(userId, newPassword)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserStore) UpdateUpdateAt(userId string) (int64, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.UpdateUpdateAt")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.UpdateUpdateAt(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserStore) VerifyEmail(userId string, email string) (string, *model.AppError) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserStore.VerifyEmail")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserStore.VerifyEmail(userId, email)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) Delete(tokenId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserAccessTokenStore.Delete(tokenId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserAccessTokenStore) DeleteAllForUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.DeleteAllForUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserAccessTokenStore.DeleteAllForUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserAccessTokenStore) Get(tokenId string) (*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.Get")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.Get(tokenId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) GetAll(offset int, limit int) ([]*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.GetAll")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.GetAll(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) GetByToken(tokenString string) (*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.GetByToken")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.GetByToken(tokenString)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) GetByUser(userId string, page int, perPage int) ([]*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.GetByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.GetByUser(userId, page, perPage)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) Save(token *model.UserAccessToken) (*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.Save(token)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) Search(term string) ([]*model.UserAccessToken, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.Search")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserAccessTokenStore.Search(term)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserAccessTokenStore) UpdateTokenDisable(tokenId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.UpdateTokenDisable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserAccessTokenStore.UpdateTokenDisable(tokenId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserAccessTokenStore) UpdateTokenEnable(tokenId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserAccessTokenStore.UpdateTokenEnable")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserAccessTokenStore.UpdateTokenEnable(tokenId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserTermsOfServiceStore.Delete")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.UserTermsOfServiceStore.Delete(userId, termsOfServiceId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerUserTermsOfServiceStore) GetByUser(userId string) (*model.UserTermsOfService, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserTermsOfServiceStore.GetByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserTermsOfServiceStore.GetByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerUserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) (*model.UserTermsOfService, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "UserTermsOfServiceStore.Save")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.UserTermsOfServiceStore.Save(userTermsOfService)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) AnalyticsIncomingCount(teamId string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.AnalyticsIncomingCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.AnalyticsIncomingCount(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.AnalyticsOutgoingCount")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.AnalyticsOutgoingCount(teamId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) ClearCaches() {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.ClearCaches")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.WebhookStore.ClearCaches()

}

func (s *OpenTracingLayerWebhookStore) DeleteIncoming(webhookId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.DeleteIncoming")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.DeleteIncoming(webhookId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) DeleteOutgoing(webhookId string, time int64) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.DeleteOutgoing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.DeleteOutgoing(webhookId, time)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncoming")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncoming(id, allowFromCache)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncomingByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncomingByChannel(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncomingByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncomingByTeam(teamId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetIncomingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncomingByTeamByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncomingByTeamByUser(teamId, userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncomingList")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncomingList(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetIncomingListByUser(userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetIncomingListByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetIncomingListByUser(userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoing(id)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingByChannel(channelId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingByChannelByUser(channelId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingByChannelByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingByChannelByUser(channelId, userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingByTeam")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingByTeam(teamId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingByTeamByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingByTeamByUser(teamId, userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingList")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingList(offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) GetOutgoingListByUser(userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.GetOutgoingListByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.GetOutgoingListByUser(userId, offset, limit)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) InvalidateWebhookCache(webhook string) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.InvalidateWebhookCache")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	s.WebhookStore.InvalidateWebhookCache(webhook)

}

func (s *OpenTracingLayerWebhookStore) PermanentDeleteIncomingByChannel(channelId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.PermanentDeleteIncomingByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.PermanentDeleteIncomingByChannel(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) PermanentDeleteIncomingByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.PermanentDeleteIncomingByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.PermanentDeleteIncomingByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.PermanentDeleteOutgoingByChannel")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.PermanentDeleteOutgoingByChannel(channelId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) PermanentDeleteOutgoingByUser(userId string) error {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.PermanentDeleteOutgoingByUser")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	err := s.WebhookStore.PermanentDeleteOutgoingByUser(userId)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return err
}

func (s *OpenTracingLayerWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.SaveIncoming")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.SaveIncoming(webhook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.SaveOutgoing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.SaveOutgoing(webhook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.UpdateIncoming")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.UpdateIncoming(webhook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayerWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	origCtx := s.Root.Store.Context()
	span, newCtx := tracing.StartSpanWithParentByContext(s.Root.Store.Context(), "WebhookStore.UpdateOutgoing")
	s.Root.Store.SetContext(newCtx)
	defer func() {
		s.Root.Store.SetContext(origCtx)
	}()

	defer span.Finish()
	result, err := s.WebhookStore.UpdateOutgoing(hook)
	if err != nil {
		span.LogFields(spanlog.Error(err))
		ext.Error.Set(span, true)
	}

	return result, err
}

func (s *OpenTracingLayer) Close() {
	s.Store.Close()
}

func (s *OpenTracingLayer) DropAllTables() {
	s.Store.DropAllTables()
}

func (s *OpenTracingLayer) GetCurrentSchemaVersion() string {
	return s.Store.GetCurrentSchemaVersion()
}

func (s *OpenTracingLayer) LockToMaster() {
	s.Store.LockToMaster()
}

func (s *OpenTracingLayer) MarkSystemRanUnitTests() {
	s.Store.MarkSystemRanUnitTests()
}

func (s *OpenTracingLayer) SetContext(context context.Context) {
	s.Store.SetContext(context)
}

func (s *OpenTracingLayer) TotalMasterDbConnections() int {
	return s.Store.TotalMasterDbConnections()
}

func (s *OpenTracingLayer) TotalReadDbConnections() int {
	return s.Store.TotalReadDbConnections()
}

func (s *OpenTracingLayer) TotalSearchDbConnections() int {
	return s.Store.TotalSearchDbConnections()
}

func (s *OpenTracingLayer) UnlockFromMaster() {
	s.Store.UnlockFromMaster()
}

func New(childStore store.Store, ctx context.Context) *OpenTracingLayer {
	newStore := OpenTracingLayer{
		Store: childStore,
	}

	newStore.AuditStore = &OpenTracingLayerAuditStore{AuditStore: childStore.Audit(), Root: &newStore}
	newStore.BotStore = &OpenTracingLayerBotStore{BotStore: childStore.Bot(), Root: &newStore}
	newStore.ChannelStore = &OpenTracingLayerChannelStore{ChannelStore: childStore.Channel(), Root: &newStore}
	newStore.ChannelMemberHistoryStore = &OpenTracingLayerChannelMemberHistoryStore{ChannelMemberHistoryStore: childStore.ChannelMemberHistory(), Root: &newStore}
	newStore.ClusterDiscoveryStore = &OpenTracingLayerClusterDiscoveryStore{ClusterDiscoveryStore: childStore.ClusterDiscovery(), Root: &newStore}
	newStore.CommandStore = &OpenTracingLayerCommandStore{CommandStore: childStore.Command(), Root: &newStore}
	newStore.CommandWebhookStore = &OpenTracingLayerCommandWebhookStore{CommandWebhookStore: childStore.CommandWebhook(), Root: &newStore}
	newStore.ComplianceStore = &OpenTracingLayerComplianceStore{ComplianceStore: childStore.Compliance(), Root: &newStore}
	newStore.EmojiStore = &OpenTracingLayerEmojiStore{EmojiStore: childStore.Emoji(), Root: &newStore}
	newStore.FileInfoStore = &OpenTracingLayerFileInfoStore{FileInfoStore: childStore.FileInfo(), Root: &newStore}
	newStore.GroupStore = &OpenTracingLayerGroupStore{GroupStore: childStore.Group(), Root: &newStore}
	newStore.JobStore = &OpenTracingLayerJobStore{JobStore: childStore.Job(), Root: &newStore}
	newStore.LicenseStore = &OpenTracingLayerLicenseStore{LicenseStore: childStore.License(), Root: &newStore}
	newStore.LinkMetadataStore = &OpenTracingLayerLinkMetadataStore{LinkMetadataStore: childStore.LinkMetadata(), Root: &newStore}
	newStore.OAuthStore = &OpenTracingLayerOAuthStore{OAuthStore: childStore.OAuth(), Root: &newStore}
	newStore.PluginStore = &OpenTracingLayerPluginStore{PluginStore: childStore.Plugin(), Root: &newStore}
	newStore.PostStore = &OpenTracingLayerPostStore{PostStore: childStore.Post(), Root: &newStore}
	newStore.PreferenceStore = &OpenTracingLayerPreferenceStore{PreferenceStore: childStore.Preference(), Root: &newStore}
	newStore.ReactionStore = &OpenTracingLayerReactionStore{ReactionStore: childStore.Reaction(), Root: &newStore}
	newStore.RoleStore = &OpenTracingLayerRoleStore{RoleStore: childStore.Role(), Root: &newStore}
	newStore.SchemeStore = &OpenTracingLayerSchemeStore{SchemeStore: childStore.Scheme(), Root: &newStore}
	newStore.SessionStore = &OpenTracingLayerSessionStore{SessionStore: childStore.Session(), Root: &newStore}
	newStore.StatusStore = &OpenTracingLayerStatusStore{StatusStore: childStore.Status(), Root: &newStore}
	newStore.SystemStore = &OpenTracingLayerSystemStore{SystemStore: childStore.System(), Root: &newStore}
	newStore.TeamStore = &OpenTracingLayerTeamStore{TeamStore: childStore.Team(), Root: &newStore}
	newStore.TermsOfServiceStore = &OpenTracingLayerTermsOfServiceStore{TermsOfServiceStore: childStore.TermsOfService(), Root: &newStore}
	newStore.TokenStore = &OpenTracingLayerTokenStore{TokenStore: childStore.Token(), Root: &newStore}
	newStore.UserStore = &OpenTracingLayerUserStore{UserStore: childStore.User(), Root: &newStore}
	newStore.UserAccessTokenStore = &OpenTracingLayerUserAccessTokenStore{UserAccessTokenStore: childStore.UserAccessToken(), Root: &newStore}
	newStore.UserTermsOfServiceStore = &OpenTracingLayerUserTermsOfServiceStore{UserTermsOfServiceStore: childStore.UserTermsOfService(), Root: &newStore}
	newStore.WebhookStore = &OpenTracingLayerWebhookStore{WebhookStore: childStore.Webhook(), Root: &newStore}
	return &newStore
}
