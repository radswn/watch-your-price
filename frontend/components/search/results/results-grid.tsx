import { SearchedObject } from 'api';
import { useEffectInit } from 'hooks/effects-lib';
import { useState } from 'react';
import { combineLatest } from 'rxjs';
import { map } from 'rxjs/operators';
import { unpackImg } from './unpack-img';
import { useMobileDetection } from 'hooks/mobile-detection';
import { Desktop } from './desktop/desktop';
import { Mobile } from './mobile/mobile';
import styles from 'styles/search.module.css';

type Args = { data: SearchedObject[] };

export function ResultsGrid(args: Args) {
    const { mobile } = useMobileDetection();
    const [preparedData, setPreparedData] = useState<SearchedObject[] | null>(null);

    useEffectInit(() => {
        const dataObservables = args.data.map(v => {
            return unpackImg(v)
                .pipe(
                    map(img => ({ ...v, img }))
                );
        });

        combineLatest(dataObservables).subscribe(data => setPreparedData(data));
    });

    if (preparedData) {
        const datView = mobile ? <Mobile data={preparedData}></Mobile> : <Desktop data={preparedData}></Desktop>;

        return <div className="grid">
            <div className="p-offset-2 p-col-8">
                <div className={`grid p-dir-col ${styles['grid_padding_bottom']}`}>
                    {datView}
                </div>
            </div>
        </div>;
    } else {
        return null;
    };
}