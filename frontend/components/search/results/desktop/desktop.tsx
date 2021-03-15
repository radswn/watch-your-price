import { SearchedObject } from 'api';
import { Button } from 'primereact/button';
import styles from 'styles/search.module.css';

type Args = {data: SearchedObject[]};

export function Desktop({data}: Args) {
    const buttons = <>
        <div>
            <Button label="Obserwuj" aria-label="Obserwuj" className="p-mx-1" />
        </div>
        <div>
            <Button label="Zobacz" aria-label="Zobacz" className="p-mx-1" />
        </div>
    </>;

    return <>
        {data!.map((v, index) => {
            return <div className="p-col p-grid p-shadow-6" key={index}>
                <div dangerouslySetInnerHTML={{ __html: v.img }} className={`p-col-4 ${styles.img}`} />
                <div className="p-col p-d-flex p-flex-column">
                    <h1 className="p-mb-2">{v.title}</h1>
                    <div className="p-mb-2">{v.description}</div>
                </div>
                <div className="p-col p-d-flex p-align-end p-justify-even">
                    {buttons}
                </div>
            </div>;
        })}
    </>;
}