# ADR (Architecture Decision Record)

ADR (Architecture Decision Record)は、システム設計における重要な決定を記録するドキュメント手法です。
プロジェクトの設計決定や変更を文書化し、その理由や背景を明確にすることが目的です。

## 作成の基準

ADRは以下のような状況で特に有用です

* 重要なアーキテクチャ上の決定が必要な場合
* 複数の選択肢があり、その中から一つを選ぶ必要がある場合
* 長期間にわたりプロジェクトに影響を与える技術的決定を行う場合

## 流れ

* **決定の文脈**: 問題の背景や、決定が必要になった状況を記述します。
* **検討した選択肢**: 考慮されたすべての選択肢と、それぞれの長所と短所を列挙します。
* **決定**: 選ばれた解決策と、その理由を詳述します。
* **結果**: 決定がもたらす期待される影響や結果について記述します。

## 更新

ADRは、新しい情報が入手されたり、状況が変化したりした場合に更新されます。
過去の決定が変更された場合は、新しいADRを作成して、変更の理由と影響を記録します。

## 初め方

以下のコマンドを実行して、テンプレートから自動的にADRを生成します。

```bash
> make gen-adr TARGET={directory path} TITLE={kebab case title of document}

ex.
> make gen-adr TARGET=apps/system/docs/tech/adr TITLE=title-of-document
```

## References

* https://github.com/joelparkerhenderson/architecture-decision-record
* https://www.ozimmer.ch/practices/2023/04/03/ADRCreation.html
