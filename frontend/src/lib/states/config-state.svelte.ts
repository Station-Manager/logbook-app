import { types } from '$lib/wailsjs/go/models';

export interface ConfigState {
    logbook: types.Logbook;
    load(this: ConfigState, cfg: types.UiConfig): void;
}

export const configState: ConfigState = $state({
    logbook: new types.Logbook(),

    load(this: ConfigState, cfg: types.UiConfig): void {
        this.logbook = cfg.logbook;
    },
});
