import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { map } from 'rxjs/operators';
import { useEffectInit } from './effects-lib';

export function useMobileDetection(mobileMaxWidth = 768) {
    const [mobile, setMobile] = useState<boolean|null>(null);
    const isMobile = window => window.screen.width <= mobileMaxWidth;
    
    //check if mobile
    useEffectInit(() => {
        setMobile(isMobile(window));

        const s = fromEvent(window, 'resize')
        .pipe(
            map((e: any) => e.target.window)
        ).subscribe(window => setMobile(isMobile(window)));

        return () => s.unsubscribe()
    });

    return {mobile};
}