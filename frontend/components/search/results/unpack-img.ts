import { SearchedObject } from "api";
import { fromEvent } from "rxjs";
import { map } from "rxjs/operators";

export function unpackImg(data: SearchedObject) {
    const buffer = Buffer.from(data!.img, 'base64');
    const blob = new Blob([buffer]);

    const fr = new FileReader();
    const o = fromEvent(fr, 'load')
        .pipe(
            map((ev: any) => ev.target!.result as string),
            map(dataURL => {
                const img = new Image();
                img.src = dataURL;
                return img.outerHTML;
            })
        )

    fr.readAsDataURL(blob);

    return o;
}