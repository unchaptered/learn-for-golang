import json
from typing import Tuple, Optional, TypedDict, List
from auto_pr_hooks import runner, splitRunner


class GitHubPr(TypedDict):
    number: int
    title: str


class GitHubPrTemplate(TypedDict):
    title: str
    body: str
    labelList: List[str]
    assigneeList: List[str]


def getPullRequestTemplate(
    *,
    base: str = 'main',
    head: str = 'dev'
):

    # [A] PR Commiter 생성
    splitedCommand = [
        'git', 'log',
        "--pretty=format:%an",
        f'{base}..{head}'
    ]
    print('[getPullRequestTemplate][Commiter] : ', splitedCommand)
    isSuccess, out, err = splitRunner(splitedCommand=splitedCommand)
    print('[getPullRequestTemplate][Commiter] isSuccess : ', isSuccess)
    print('[getPullRequestTemplate][Commiter] out : ', out)
    print('[getPullRequestTemplate][Commiter] err : ', err)
    commiterList = list(set(out.split('\n')))

    # [B] PR Body 생성
    splitedCommand = [
        'git', 'log',
        f'--pretty=format:"- `%ad` %h **%s** %an"',
        f'--date=format-local:%y-%m-%d %H:%M',
        f'{base}..{head}'
    ]
    print('[getPullRequestTemplate][Logs] : ', splitedCommand)
    isSuccess, out, err = splitRunner(splitedCommand=splitedCommand)
    print('[getPullRequestTemplate][Logs] isSuccess : ', isSuccess)
    print('[getPullRequestTemplate][Logs] out : ', out)
    print('[getPullRequestTemplate][Logs] err : ', err)

    splitedCommandToGetRepoData = ['git', 'remote', 'get-url', 'origin']
    isSuccess2, out2, err2 = splitRunner(
        splitedCommand=splitedCommandToGetRepoData)
    owner: str = None
    reponame: str = None

    if isSuccess2:
        print('isSuccess2 : ', isSuccess2)
        print('out2 : ', out2)
        print('err2 : ', err2)
        splitedOutput = out2.split('/')

        owner: str = splitedOutput[3]
        reponame: str = splitedOutput[4]
        reponame = reponame.replace('\n', '').replace('.git', '')

    pullRequestBody = f"""
[제목] 제목 변경 필요
[담당] {','.join(commiterList)}
[프로젝트] {owner} / {reponame}
[내용] 작성 필요
[반영정보]
"""
    commitList = out.split('\n')
    for commit in commitList:
        pullRequestBody += '\n' + commit[1:-1]
    print("✅✅✅")
    print(pullRequestBody)
    # PR Label 생성
    splitedCommand = [
        'git', 'log',
        f'--pretty=format:%s',
        f'--date=format-local:%y-%m-%d %H:%M',
        f'{base}..{head}'
    ]
    isSuccess, out, err = splitRunner(splitedCommand=splitedCommand)
    commitList = out.split('\n')

    tokenKeys = ['create', 'update', 'fix']
    tokenValues = ['enhancement', 'enhancement', 'bug']
    labelList = []
    for commit in commitList:
        for idx, tokenKey in enumerate(tokenKeys):
            hasToken = tokenKey in commit.lower()
            if hasToken:
                labelList.append(tokenValues[idx])
    labelList = list(set(labelList))

    pullRequestBody += """
    
[문제] 없음
[기타]
    """

    githubPrTemplate: GitHubPrTemplate = {
        'title': '제목 변경 필요',
        'body': pullRequestBody,
        'labelList': labelList,
        'assigneeList': commiterList
    }

    return githubPrTemplate


def fetchGitHub(
    *,
    base: str = 'main',
    head: str = 'dev'
):
    command = 'git fetch origin main:main'
    isSuccess, outStr, errStr = runner(command=command)
    if not isSuccess:
        raise SyntaxError([
            'command syntax occure this error',
            errStr
        ])


def isExistsPrList(
    *,
    base: str = 'main',
    head: str = 'dev',
    jsonFormat: Optional[str] = None
) -> Tuple[bool, GitHubPr]:
    command = f'gh pr list --base {base} --head {head}'
    if jsonFormat:
        command += ' --json number,title'

    isSuccess, outStr, errStr = runner(command=command)
    if not isSuccess:
        raise SyntaxError([
            'command syntax occure this error',
            errStr
        ])

    prList: List[GitHubPr] = json.loads(outStr)
    hasPr = len(prList) != 0
    return hasPr, prList


def updatePullRequest(
    *,
    prList: List[GitHubPr],
    githubPrTemplate: GitHubPrTemplate
):
    pr = prList[0]

    title = githubPrTemplate['title']
    body = githubPrTemplate['body']
    labelList = githubPrTemplate['labelList']
    assigneeList = githubPrTemplate['assigneeList']

    splitedCommand = [
        'gh', 'pr', 'edit', str(pr['number']),
        '--title', str(title),
        '--body', f'{body}',
    ]

    TK = ','
    hasLabel = len(labelList) > 0
    if hasLabel:
        label = TK.join(labelList)
        splitedCommand.extend(['--add-label', label])

    hasAssignee = len(assigneeList) > 0
    if hasAssignee:
        assignee = TK.join(assigneeList)
        splitedCommand.extend(['--add-assignee', assignee])

    print('[updatePullRequest] splitedCommand : ', splitedCommand)
    isSuccess, outStr, errStr = splitRunner(splitedCommand=splitedCommand)
    print('[updatePullRequest] isSuccess : ', isSuccess)
    print('[updatePullRequest] outStr : ', outStr)
    print('[updatePullRequest] errStr : ', errStr)


def createPullRequest(
    *,
    base: str,
    head: str,
    githubPrTemplate: GitHubPrTemplate
):

    title = githubPrTemplate['title']
    body = githubPrTemplate['body']
    labelList = githubPrTemplate['labelList']
    assigneeList = githubPrTemplate['assigneeList']

    # command = f'gh pr create --base {base} --head {head} --title {title} --body {body}'
    splitedCommand = [
        'gh', 'pr', 'create',
        '--base', base,
        '--head', head,
        '--title', title,
        '--body', body
    ]

    TK = ','
    hasLabel = len(labelList) > 0
    if hasLabel:
        label = TK.join(labelList)
        splitedCommand.extend(['--label', label])

    hasAssignee = len(assigneeList) > 0
    if hasAssignee:
        assignee = TK.join(assigneeList)
        splitedCommand.extend(['--assignee', assignee])

    print('[createPullRequest] splitedCommand : ', splitedCommand)
    isSuccess, outStr, errStr = splitRunner(splitedCommand=splitedCommand)
    print('[createPullRequest] isSuccess : ', isSuccess)
    print('[createPullRequest] outStr : ', outStr)
    print('[createPullRequest] errStr : ', errStr)


if __name__ == '__main__':
    import sys

    if len(sys.argv) == 1:
        raise ValueError(
            f'python ./scripts/{sys.argv[0]} <base> <head>에서 base가 누락되었습니다.')
    if len(sys.argv) == 2:
        raise ValueError(
            f'python ./scripts/{sys.argv[0]} <base> <head>에서 head가 누락되었습니다.')

    base = sys.argv[1]
    head = sys.argv[2]
    # fetchGitHub(base=base,
    #             head=head)
    hasPr, prList = isExistsPrList(base=base,
                                   head=head,
                                   jsonFormat='number,title')
    requestTemplate = getPullRequestTemplate(base=base,
                                             head=head)

    if hasPr:
        updatePullRequest(prList=prList,
                          githubPrTemplate=requestTemplate)

    else:
        createPullRequest(base=base,
                          head=head,
                          githubPrTemplate=requestTemplate)
