import { InputText } from 'primereact/inputtext';
import { useEffect, useState } from 'react';
import styles from 'styles/search.module.css';

type Args = {onChange: (string) => any};

export function DesktopSearch(args: Args) {
    const [value, setValue] = useState<string>('');

    useEffect(() => args.onChange(value), [value])

    return <>
        <div className="grid">
            <div className="p-col-8 p-offset-2">
                <span className={`p-input-icon-left ${styles['desktop-search']}`}>
                    <i className="pi pi-search" />
                    <InputText value={value} onChange={(e: any) => setValue(e.target.value)} 
                    placeholder="Wyszukaj" className={styles['desktop-search']}/>
                </span>
            </div>
        </div>
    </>;
}