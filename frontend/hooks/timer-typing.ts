import { useEffectUpdateNullish } from "hooks/effects-lib";
import { useState } from "react";
import { Subscription, timer } from "rxjs";

export function useTimerTyping(text: string|null) {
    const [timeoutTyping, setTimeoutTyping] = useState<boolean>(false);
    const [subscription, setSubscription] = useState<Subscription|null>(null);

    useEffectUpdateNullish(() => {
        setTimeoutTyping(false);
        subscription?.unsubscribe();
        
        const s = timer(1000)
        .subscribe(() => setTimeoutTyping(true));

        setSubscription(s);

        return () => s.unsubscribe();
    }, [text]);

    return {timeoutTyping};
}