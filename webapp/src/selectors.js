import {id as pluginId} from './manifest';

const getPluginState = (state) => state['plugins-' + pluginId] || {};

export const getMeetingSettingsModalState = (state) => getPluginState(state).meetingSettingsModal;
export const getMeetingSettings = (state) => getPluginState(state).meetingSettings;
