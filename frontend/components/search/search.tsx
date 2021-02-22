import { MobileSearch } from './mobile/mobile';
import { DesktopSearch } from './desktop/desktop';
import { useMobileDetection } from 'hooks/mobile-detection';

export function Search() {
    const {mobile} = useMobileDetection();

    return mobile ? <MobileSearch onChange={v => {}}/> : <DesktopSearch onChange={v => {}}/>;
}