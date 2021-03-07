import { useMousePosition } from 'hooks/mouse-position';
import { SidebarAnimation } from './sidebar-animation';
import { useState } from 'react';
import { Sidebar } from 'primereact/sidebar';
import styles from 'styles/navbar.module.css';
import { useEffectInit, useEffectUpdate } from 'hooks/effects-lib';

export function MovableLeftSidebar({content, className = ''}) {
    const [marginWidthPx, fadeInInitListenerWidth] = [100, 100];
    const [animation, setAnimation] = useState<SidebarAnimation|null>(null);
    const {mousePos, rightMoveDirection, moving} = useMousePosition({fadeInInitListenerWidth, fadeInMaxWidth: animation?.sidebar.width ?? null});

    useEffectInit(() => {
        const sidebar = document.querySelector('.p-sidebar-left') as HTMLElement;
        sidebar.style.left = `-${sidebar.clientWidth + marginWidthPx}px`;

        const blur = document.querySelector('#blur') as HTMLElement;

        const sidebarInfo = {width: sidebar.clientWidth, component: sidebar};
        setAnimation(new SidebarAnimation(sidebarInfo, blur, fadeInInitListenerWidth, marginWidthPx));
    });

    useEffectUpdate(() => {
        animation?.animate(mousePos!, rightMoveDirection!, moving!);
    }, [mousePos, rightMoveDirection, moving]);

    return <>
        <Sidebar visible={true} onHide={() => {}} showCloseIcon={false} className={`${className} ${styles['movable-sidebar']}`}>
            {content}
        </Sidebar>

        <div id="blur" className={styles['blur']}></div>
    </>;
}