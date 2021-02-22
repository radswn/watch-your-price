import { MobileNavbar } from './mobile/mobile';
import { DesktopNavbar } from './desktop/desktop';
import { useMobileDetection } from 'hooks/mobile-detection';

export function Navbar() {
    const {mobile} = useMobileDetection();

    return mobile ? <MobileNavbar/> : <DesktopNavbar/>;
}