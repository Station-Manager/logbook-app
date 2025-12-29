import { FetchUiConfig } from '$lib/wailsjs/go/facade/Service';
import { LogError } from '$lib/wailsjs/runtime';
import { types } from '$lib/wailsjs/go/models';
import { configState } from '$lib/states/config-state.svelte';
import { handleAsyncError } from '$lib/utils/error-handler';

export const prerender = true;
export const ssr = false;

export const load = async (): Promise<void> => {
    try {
        const cfg: types.UiConfig | null | undefined = await FetchUiConfig();
        let activeCfg = cfg;
        if (!activeCfg) {
            const msg = 'UiConfig fetch returned null or undefined; using defaults.';
            LogError(msg);
            activeCfg = new types.UiConfig();
        }
        // Load the configuration into the state object
        configState.load(activeCfg);
    } catch (e: unknown) {
        handleAsyncError(e, '+layout.ts->load: LayoutData');
    }
};
