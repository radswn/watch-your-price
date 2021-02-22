import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { map } from 'rxjs/operators';

export function useMobileDetection(mobileMaxWidth = 768): {mobile: boolean|null} {
    const [mobile, setMobile] = useState<boolean|null>(null);
    const isMobile = window => window.screen.width <= mobileMaxWidth;
    
    //check if mobile
    useEffect(() => {
        setMobile(isMobile(window));

        const s = fromEvent(window, 'resize')
        .pipe(
            map((e: any) => e.target.window)
        ).subscribe(window => setMobile(isMobile(window)));

        return () => s.unsubscribe()
    }, []);

    return {mobile};
}