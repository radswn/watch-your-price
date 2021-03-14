import { useState } from 'react';
import { mergeMap } from 'rxjs/operators';
import { SearchedObject } from 'api';
import { fromFetch } from 'rxjs/fetch'
import { useEffectUpdateNullish } from './effects-lib';

export function useFetchSearch(text: string|null) {
    const [data, setData] = useState<SearchedObject[]|null>(null);
    const [error, setError] = useState<Error|null>(null);
    
    useEffectUpdateNullish(() => {
        fromFetch(`http://localhost:3000/api/search?item=${text}`)
        .pipe(
            mergeMap((r: Response) => r.json())
        )
        .subscribe({
            error: (e: Error) => setError(e),
            next: (d: SearchedObject[]) => setData(d)
        });
    }, [text]);

    return {data, error};
}