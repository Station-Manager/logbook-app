import type { PageData } from './$types';
import { GetLogbookList } from '$lib/wailsjs/go/facade/Service';
import { handleAsyncError } from '$lib/utils/error-handler';
import { types } from '$lib/wailsjs/go/models';

export const load: PageData = async (): Promise<object> => {
    try {
        const list: types.Logbook[] = await GetLogbookList();
        return { list: list };
    } catch (e: unknown) {
        handleAsyncError(e, 'database/+page.ts->load');
    }
    return { list: null };
};
