// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

/* eslint-disable no-console */
/* eslint-disable global-require */
/* eslint-disable func-names */
/* eslint-disable prefer-arrow-callback */
/* eslint-disable no-magic-numbers */
/* eslint-disable no-unreachable */

import assert from 'assert';
import TestHelper from './test_helper.jsx';

describe('Client.OAuth', function() {
    this.timeout(100000);

    // it('registerOAuthApp', function(done) {
    //     TestHelper.initBasic(() => {
    //         TestHelper.basicClient().enableLogErrorsToConsole(false); // Disabling sicne this unit test causes an error

    //         var app = {};
    //         app.name = 'test';
    //         app.homepage = 'homepage';
    //         app.description = 'desc';
    //         app.callback_urls = '';

    //         TestHelper.basicClient().registerOAuthApp(
    //             app,
    //             function() {
    //                 done(new Error('not enabled'));
    //             },
    //             function(err) {
    //                 assert.equal(err.id, 'api.oauth.register_oauth_app.turn_off.app_error');
    //                 done();
    //             }
    //         );
    //     });
    // });

    // it('allowOAuth2', function(done) {
    //     TestHelper.initBasic(() => {
    //         TestHelper.basicClient().enableLogErrorsToConsole(false); // Disabling sicne this unit test causes an error

    //         TestHelper.basicClient().allowOAuth2(
    //             'GET',
    //             '123456',
    //             'http://nowhere.com',
    //             'state',
    //             'scope',
    //             function() {
    //                 done(new Error('not enabled'));
    //             },
    //             function(err) {
    //                 assert.equal(err.id, 'api.oauth.allow_oauth.turn_off.app_error');
    //                 done();
    //             }
    //         );
    //     });
    // });
});
