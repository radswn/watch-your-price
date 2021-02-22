type SidebarInfo = {
    width: number|null, 
    component: HTMLElement|null
};

export class SidebarAnimation {
    private sidebarFrames: {left}[] = [];
    private blurFrames: {background, opacity}[] = [];

    constructor(
        public readonly sidebar: SidebarInfo, 
        public readonly blur: HTMLElement, 
        public readonly fadeInInitListenerWidth: number,
        public readonly marginWidthPx: number) { }

    animate(mousePos: number, rightMoveDirection: boolean, moving: boolean) {
        if(rightMoveDirection) {
            if(mousePos <= this.fadeInInitListenerWidth && moving) {
                this.moveToClickpoint(mousePos);
            } else if(mousePos! <= this.fadeInInitListenerWidth) {
                this.setHide();
            } else {
                this.setFullWidth();
            }
        } else {
            if(mousePos <= this.fadeInInitListenerWidth) {
                this.setHide();
            } else if(moving) {
                this.moveToClickpoint(mousePos);
            } else {
                this.setFullWidth();
            }
        }

        this.sidebar.component?.animate(this.sidebarFrames, {
            duration: 1000, 
            fill: 'forwards'
        });
        this.blur?.animate(this.blurFrames, {
            duration: 1000, 
            fill: 'forwards'
        })

        this.sidebarFrames.splice(0);
        this.blurFrames.splice(0);
    }

    private setOpaque(value: number) {
        if(this.blur) {
            //first number decreases difference between contrasts colors
            const opacity = 1 - value;
            const light = 8*value * 100;
            
            this.blurFrames.push({
                background: `hsl(0, 0%, ${light}%)`, 
                opacity
            });
        }
    };

    private setHide() {
        this.sidebarFrames.push({left: `-${this.marginWidthPx + this.sidebar.width!}px`});
        this.setOpaque(1);
    };

    private setFullWidth() {
        this.sidebarFrames.push({left: '0px'})
        const opaquqe = (window.screen.width - this.sidebar.width!)/window.screen.width;
        this.setOpaque(opaquqe);
    };

    private moveToClickpoint(mousePos: number) {
        this.sidebarFrames.push({left: `${mousePos - this.sidebar.width!}px`});
        const opaque = (window.screen.width - mousePos)/window.screen.width;
        this.setOpaque(opaque);
    };
}