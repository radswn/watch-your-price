import { Toolbar } from 'primereact/toolbar';
import { Button } from 'primereact/button';
import styles from 'styles/navbar.module.css';

export function DesktopNavbar() {
    const left = <>
        <Button label="Wyszukiwarka" aria-label="Wyszukiwarka" className="p-button-inverse"/>
        <Button label="Historia cen" aria-label="Historia cen" className="p-button-inverse"/>
        <Button label="Użytkownik" aria-label="Użytkownik" className="p-button-inverse"/>
    </>;

    const right = <Button label="Wyloguj" aria-label="Wyloguj" className="p-button-inverse"/>;

    return <Toolbar left={() => left} right={() => right} className={styles['nav-color']}/>
}