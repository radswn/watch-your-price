import { Button } from 'primereact/button';
import { MovableLeftSidebar } from './movable-sidebar/movable-sidebar';
import styles from 'styles/navbar.module.css';

export function MobileNavbar() {
    const content = <>
        <h1 className={styles['header']}>Akcje</h1>
        <div className="p-grid p-dir-col">
            <Button label="Wyloguj" aria-label="Logout" className="p-my-1 p-col p-button-inverse p-shadow-5"/>
            <Button label="Wyszukiwarka" aria-label="Wyszukiwarka" className="p-my-1 p-col p-button-inverse p-shadow-5"/>
            <Button label="Historia cen" aria-label="Historia cen" className="p-my-1 p-col p-button-inverse p-shadow-5"/>
            <Button label="Użytkownik" aria-label="Użytkownik" className="p-my-1 p-col p-button-inverse p-shadow-5"/>
        </div>
    </>;

    return <MovableLeftSidebar content={content} className={styles['nav-color']}></MovableLeftSidebar>
}