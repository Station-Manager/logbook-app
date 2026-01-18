import { types } from '$lib/wailsjs/go/models';

export interface ConfigState {
    logbook: types.Logbook;
    pageSize: number;
    load(this: ConfigState, cfg: types.UiConfig): void;
}

export const configState: ConfigState = $state({
    logbook: new types.Logbook(),
    pageSize: 24,
    load(this: ConfigState, cfg: types.UiConfig): void {
        this.logbook = cfg.logbook;
        this.pageSize = cfg.pagination_page_size;
    },
});
