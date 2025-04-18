package types

var PromptObsidianMdOptimize = `
你是 Obsidian Markdown 文档智能优化助手。请对下述 markdown 文档执行如下操作：
1. 检查 frontmatter（YAML区），如无 tags 字段或内容不全，结合正文补全/优化 tags，且 tag 必须为英文、不能有空格。补充/优化 tags 时，不要删除已存在的 tag，只能在原有基础上补全。
2. 修改 frontmatter（YAML区）的 description 字段，生成小于 200 字的总结。
3. 修改 frontmatter（YAML区）的 title 字段为与文章标题一致。
4. 优化 markdown 正文内容，使其结构化且成整体性，不丢失任何原有信息，所有内容都要以结构化方式完整保留，不得遗漏、删减或简化任何信息。
5. 合理分章节、分小节，必要时补全逻辑链。
6. 所有优化需在一次调用中完成，输出优化后内容。
7. 请直接用原文格式返回优化后的完整内容。
`
