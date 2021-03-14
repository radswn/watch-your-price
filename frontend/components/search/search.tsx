import { MobileSearch } from './mobile/mobile';
import { DesktopSearch } from './desktop/desktop';
import { useMobileDetection } from 'hooks/mobile-detection';
import { useState } from 'react';
import { useFetchSearch } from 'hooks/fetch-search';
import { ResultsGrid } from './results/results-grid';
import { useTimerTyping } from 'hooks/timer-typing';
import { useEffectUpdateValue } from 'hooks/effects-lib';

export function Search() {
    const {mobile} = useMobileDetection();
    const [text, setText] = useState<string|null>(null);
    const {timeoutTyping} = useTimerTyping(text);
    const [textToSearch, setTextToSearch] = useState<string|null>(null);
    const {data, error} = useFetchSearch(textToSearch);

    useEffectUpdateValue(() => setTextToSearch(text), [{dep: timeoutTyping, vals: [true]}]);

    const search = mobile ? <MobileSearch onChange={v => setText(v)}/> : <DesktopSearch onChange={v => setText(v)}/>;
    const results = data ? <ResultsGrid data={data}/> : null

    return <>
        {search}
        {results}
        {error}
    </>;
}