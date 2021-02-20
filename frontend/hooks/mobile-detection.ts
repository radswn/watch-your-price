import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { map } from 'rxjs/operators';

type Inputs = {mobileMaxWidth: number};
type Returns = {mobile: boolean|null}

export function useMobileDetection(args: Inputs): Returns {
    const [mobile, setMobile] = useState<boolean|null>(null);
    const isMobile = window => window.screen.width <= args.mobileMaxWidth;
    
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