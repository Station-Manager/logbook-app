import type { PageData } from './$types';
import { GetDatabaseFileList } from '$lib/wailsjs/go/facade/Service';
import { handleAsyncError } from '$lib/utils/error-handler';

export const load: PageData = async (): Promise<object> => {
    console.log('load called');
    try {
        const list: string[] = await GetDatabaseFileList();
        console.log('Database file list:', list);
        return { list: list };
    } catch (e: unknown) {
        handleAsyncError(e, 'database/+page.ts->load');
    }
    return { list: null };
};
