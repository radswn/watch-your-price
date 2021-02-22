import { InputText } from 'primereact/inputtext';
import { Toolbar } from 'primereact/toolbar';
import { useEffect, useState } from 'react';
import styles from 'styles/search.module.css';

type Args = {onChange: (string) => any};

export function MobileSearch(args: Args) {
    const [value, setValue] = useState<string>('');

    const content = <>
        <span className={`p-input-icon-left ${styles['mobile-search']}`}>
            <i className="pi pi-search" />
            <InputText value={value} onChange={(e: any) => setValue(e.target.value)} 
            placeholder="Wyszukaj" className={styles['mobile-search']}/>
        </span> 
    </>;

    useEffect(() => args.onChange(value), [value])

    return <>
        <Toolbar left={() => content} className={styles['mobile-toolbar']}/>
        <style jsx global>{`
            .p-toolbar > *:first-child {
                width: 100% !important;
            }
        `}</style>
    </>;
}