import type { PageData } from './$types';

export const load: PageData = async (): Promise<object> => {
    console.log('load called');
    return {};
};
