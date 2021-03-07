import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { map } from 'rxjs/operators';
import { useEffectUpdate } from './effects-lib';

type Inputs = {
    fadeInInitListenerWidth: number|null, 
    fadeInMaxWidth: number|null
};

export function useMousePosition(args: Inputs) {
    type Position = {
        prev: number|null, 
        current: number|null
    };

    const [mousePos, setMousePos] = useState<Position>({prev: null, current: null});
    const [rightMoveDirection, setRightMoveDirection] = useState<boolean|null>(null);
    const [moving, setMoving] = useState<boolean|null>(null);

    useEffectUpdate(() => {
        let listen = false;
        let prev: number|null = null;

        const downEvent = fromEvent(window, 'pointerdown')
        .pipe(
            map((e: any) => e.clientX)
        ).subscribe((current: number) => {
            if(current <= args.fadeInInitListenerWidth! || prev ) {
                listen = true;
                setMoving(true);
            }
        });

        const moveEvent = fromEvent(window, 'pointermove')
        .pipe(
            map((e: any) => e.clientX),
        ).subscribe((current: number) => {
            if(current <= args.fadeInMaxWidth! && listen) {
                setMousePos({prev, current});
                prev = current;
            }
        });

        const upEvent = fromEvent(window, 'pointerup')
        .pipe(
            map((e: any) => e.clientX),
        ).subscribe((current: number) => {
            listen = false;
            setMoving(false);
            if(current <= args.fadeInInitListenerWidth!) {
                prev = null;
            }
        });


        return () => {
            downEvent.unsubscribe();
            moveEvent.unsubscribe();
            upEvent.unsubscribe();
        };
    }, [args.fadeInMaxWidth]);


    useEffectUpdate(() => {
        const {current, prev} = mousePos;
        setRightMoveDirection(current! >= prev!);
    }, [mousePos]);

    return {mousePos: mousePos.current, rightMoveDirection, moving};
}
