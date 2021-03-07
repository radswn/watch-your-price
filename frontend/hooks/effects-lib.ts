import { DependencyList, EffectCallback, useEffect } from "react";

export function useEffectInit(effect: EffectCallback) {
    useEffect(effect, []);
}

export function useEffectUpdate(effect: EffectCallback, deps: DependencyList) {
    useEffect(effect, deps);
}

export function useEffectUpdateNullish(effect: EffectCallback, deps: DependencyList) {
    useEffect(() => {
        for(const d of deps) {
            if(!d) {
                return;
            }
        }

        return effect();
    }, deps);
}