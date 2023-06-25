import { _decorator, Component, Node } from 'cc';

export class GameModel extends Component {
    private static instance: GameModel | null = null;

    public static get Instance(): GameModel {
        return GameModel.instance!;
    }

    onLoad(): void {
        if(!GameModel.instance) {
            GameModel.instance = this;
        } else {
            this.destroy();
            return;
        }
    }
}