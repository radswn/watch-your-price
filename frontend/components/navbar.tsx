import { Toolbar } from 'primereact/toolbar';
import { Sidebar } from 'primereact/sidebar';
import { Button } from 'primereact/button';
import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { map } from 'rxjs/operators';
import styles from 'styles/navbar.module.css';

function useMobileDetection(): boolean|null {
    const [mobile, setMobile] = useState<boolean|null>(null);
    const isMobile = window => window.screen.width < 769;
    
    useEffect(() => {
        setMobile(isMobile(window));

        const s = fromEvent(window, 'resize')
        .subscribe((e: any) => setMobile(isMobile(e.target.window)));

        return () => s.unsubscribe()
    }, []);

    return mobile;
}

function useMousePosition(maxWidth: number|null): [number|null, boolean|null] {
    type Position = {
        prev: number|null, 
        current: number|null
    };

    const [mousePos, setMousePos] = useState<Position>({prev: null, current: null});
    const [right, setRight] = useState<boolean|null>(null);

    useEffect(() => {
        let listen = false;
        let prev: number|null = null;

        const downEvent = fromEvent(window, 'pointerdown').subscribe(() => listen = true);

        const moveEvent = fromEvent(window, 'pointermove')
        .pipe(
            map((e: any) => e.clientX),
            map((current: number) => {
                const next = () => {
                    if(current <= maxWidth!) {
                        setMousePos({prev, current});
                        prev = current;
                    }
                }
                
                return listen ? next : null;
            })
        ).subscribe(callback => callback?.());

        const upEvent = fromEvent(window, 'pointerup').subscribe(() => listen = false)


        return () => {
            downEvent.unsubscribe();
            moveEvent.unsubscribe();
            upEvent.unsubscribe();
        };
    }, [maxWidth]);

    useEffect(() => {
        const {current, prev} = mousePos;
        setRight(current! >= prev!);
    }, [mousePos]);

    return [mousePos.current, right];
}

function MovableLeftSidebar({content, className = ''}) {
    type SidebarInfo = {
        width: number|null, 
        component: HTMLElement|null
    };

    const marginWidthPx = 100;
    const [sidebar, setSidebar] = useState<SidebarInfo>({width: null, component: null});
    const [mousePos, right] = useMousePosition(sidebar.width);

    useEffect(() => {
        const component = document.querySelector('.p-sidebar-left') as HTMLElement;
        component.style.left = `-${component.clientWidth + marginWidthPx}px`;
        setSidebar({width: component.clientWidth, component});
    }, []);

    useEffect(() => {
        const left = !right && mousePos! < 50 ? `-${sidebar.width! + marginWidthPx}px` : `${mousePos! - sidebar.width!}px`;

        sidebar.component?.animate([{left}], {
            duration: 1000,
            fill: 'forwards'
        })
    }, [mousePos, right])

    return <Sidebar visible={true} onHide={() => {}} showCloseIcon={false} className={className}>
        {content}
    </Sidebar>;
}

function MobileNavbar() {
    const content = <>
        <h1 style={{color: 'white'}}>Actions</h1>
        <div className="p-d-flex p-flex-column">
            <Button label="Logout" aria-label="Logout" className="p-mb-2 p-button-inverse"/>
            <Button label="Search" aria-label="Search" className="p-mb-2 p-button-inverse"/>
            <Button label="History" aria-label="History" className="p-mb-2 p-button-inverse"/>
            <Button label="Account" aria-label="Account" className="p-mb-2 p-button-inverse"/>
        </div>
    </>;

    return <MovableLeftSidebar content={content} className={styles['nav-color']}></MovableLeftSidebar>
}

function DesktopNavbar() {
    const left = <>
        <Button label="Search" aria-label="Search" className="p-button-inverse"/>
        <Button label="History" aria-label="History" className="p-button-inverse"/>
        <Button label="Account" aria-label="Account" className="p-button-inverse"/>
    </>;

    const right = <Button label="Logout" aria-label="Logout" className="p-button-inverse"/>;

    return <Toolbar left={() => left} right={() => right} className={styles['nav-color']}/>
}

export function Navbar() {
    const mobile = useMobileDetection();

    return mobile ? <MobileNavbar/> : <DesktopNavbar/>;
}