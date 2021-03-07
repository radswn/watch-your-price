import { useEffectUpdate } from 'hooks/effects-lib';
import { InputText } from 'primereact/inputtext';
import { useState } from 'react';
import styles from 'styles/search.module.css';

type Args = {onChange: (string) => any};

export function DesktopSearch(args: Args) {
    const [value, setValue] = useState<string>('');

    useEffectUpdate(() => args.onChange(value), [value])

    return <>
        <div className="grid">
            <div className={`p-offset-2 p-col-8 ${styles['desktop-search-wrapper']}`}>
                <span className={`p-input-icon-left ${styles['desktop-search']}`}>
                    <i className="pi pi-search" />
                    <InputText value={value} onChange={(e: any) => setValue(e.target.value)} 
                    placeholder="Wyszukaj" className={styles['desktop-search']}/>
                </span>
            </div>
        </div>
    </>;
}