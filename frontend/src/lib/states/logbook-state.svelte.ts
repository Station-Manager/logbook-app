import { types } from '$lib/wailsjs/go/models';

export interface LogbookState {
    fromLogbook(this: LogbookState, lb: types.Logbook): void;
    toLogbook(this: LogbookState): types.Logbook;
}

export const logbookState: LogbookState = $state({
    fromLogbook(this: LogbookState, lb: types.Logbook): void {
        console.log(lb);
    },
    toLogbook(this: LogbookState): types.Logbook {
        const base = new types.Logbook();
        console.log(base);
        return base;
    },
});
