import { SearchedObject } from 'api';
import { useEffectInit, useEffectUpdateNullish } from 'hooks/effects-lib';
import { useState } from 'react';
import { combineLatest } from 'rxjs';
import { map } from 'rxjs/operators';
import { unpackImg } from './unpack-img';
import { Button } from 'primereact/button';
import styles from 'styles/search.module.css';

type Args = { data: SearchedObject[] };

export function ResultsGrid(args: Args) {
    const [preparedData, setPreparedData] = useState<SearchedObject[] | null>(null);
    const [htmlGrid, setHtmlGrid] = useState<JSX.Element[] | null>(null);

    useEffectInit(() => {
        const dataObservables = args.data.map(v => {
            return unpackImg(v)
                .pipe(
                    map(img => ({ ...v, img }))
                );
        });

        combineLatest(dataObservables).subscribe(data => setPreparedData(data));
    });

    useEffectUpdateNullish(() => {
        const buttons = <>
            <div>
                <Button label="Obserwuj" aria-label="Obserwuj" className="p-mx-1"/>
            </div>
            <div>
                <Button label="Zobacz" aria-label="Zobacz" className="p-mx-1"/>
            </div>
        </>;

        const grid = preparedData!.map((v, index) => {
            return <div className="p-col p-grid p-shadow-6" key={index}>
                <div dangerouslySetInnerHTML={{ __html: v.img }} className={`p-col-4 ${styles.img}`} />
                <div className="p-col p-d-flex p-flex-column">
                    <div className="p-mb-2">{v.title}</div>
                    <div className="p-mb-2">{v.description}</div>
                </div>
                <div className="p-col p-d-flex p-align-end p-justify-even">
                    {buttons}
                </div>
            </div>;
        })

        setHtmlGrid(grid);
    }, [preparedData]);

    return <div className="grid">
        <div className="p-offset-2 p-col-8">
            <div className={`grid p-dir-col ${styles['grid_padding_bottom']}`}>
                {htmlGrid}
            </div>
        </div>
    </div>
}