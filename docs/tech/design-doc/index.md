# Design Doc

Design Docとはシステム設計ドキュメント手法のこと。
開発をする前にプロジェクトの背景や目的、設計、検討した代案などをドキュメント化することを目的としています。

## 作成の基準

明確な基準はないが、以下の場合にDesign Docの作成をする場合が多い

* 開発完了までに数sprintを要する
* いくつかの実装方法が考えられる
* 技術的であったりドメイン的に新規のものや慣れていないものを扱う

## 流れ

開発中に変更があった際に、作成者の思考の整理としてDesign Docsを更新することはありますが、必須ではない。
また、開発終了後に仕様変更があった場合もDesign Docsを更新することはありません。
あくまで開発時のSnapshotとして議論することを最大の目的としています。

開発完了後にはDesign DocsのstatusをArchiveに更新します。

## 初め方

以下コマンドを実行しテンプレートから自動的にDesign Docが生成されます

```bash
> make gen-design-doc TARGET={directory path} TITLE={kebab case title of document}

ex.
> make gen-design-doc TARGET=apps/system/docs/tech/design-doc TITLE=title-of-document
```

## References

* https://www.industrialempathy.com/posts/design-docs-at-google
* https://engineering.mercari.com/blog/entry/20220225-design-docs-by-mercari-shops
* https://r-kaga.com/blog/collection-of-resources-for-writing-a-good-design-docs
