// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import {intlShape, injectIntl, defineMessages} from 'react-intl';
import Constants from '../utils/constants.jsx';
import GetLinkModal from './get_link_modal.jsx';
import ModalStore from '../stores/modal_store.jsx';
import TeamStore from '../stores/team_store.jsx';

const messages = defineMessages({
    title: {
        id: 'get_team_invite_link_modal.title',
        defaultMessage: 'Team Invite Link'
    },
    help: {
        id: 'get_team_invite_link_modal.help',
        defaultMessage: 'Send teammates the link below for them to sign-up to this team site.'
    }
});

class GetTeamInviteLinkModal extends React.Component {
    constructor(props) {
        super(props);

        this.handleToggle = this.handleToggle.bind(this);

        this.state = {
            show: false
        };
    }

    componentDidMount() {
        ModalStore.addModalListener(Constants.ActionTypes.TOGGLE_GET_TEAM_INVITE_LINK_MODAL, this.handleToggle);
    }

    componentWillUnmount() {
        ModalStore.removeModalListener(Constants.ActionTypes.TOGGLE_GET_TEAM_INVITE_LINK_MODAL, this.handleToggle);
    }

    handleToggle(value) {
        this.setState({
            show: value
        });
    }

    render() {
        const {formatMessage} = this.props.intl;

        return (
            <GetLinkModal
                show={this.state.show}
                onHide={() => this.setState({show: false})}
                title={formatMessage(messages.title)}
                helpText={formatMessage(messages.help)}
                link={TeamStore.getCurrentInviteLink()}
            />
        );
    }
}

GetTeamInviteLinkModal.propTypes = {
    intl: intlShape.isRequired
};

export default injectIntl(GetTeamInviteLinkModal);