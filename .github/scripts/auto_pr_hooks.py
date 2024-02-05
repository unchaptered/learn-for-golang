import subprocess
from typing import Tuple, List


def runnerWrapper(
    *,
    splitedCommand: List[str],
    encoding: str = 'utf-8'
) -> subprocess.CompletedProcess:
    return subprocess.run(splitedCommand,
                          stdin=subprocess.PIPE,
                          stdout=subprocess.PIPE,
                          stderr=subprocess.PIPE,
                          encoding=encoding)


def runner(
    *,
    command: str
) -> Tuple[bool, str, str]:

    result = runnerWrapper(splitedCommand=command.split(' '))
    isSuccess = result.stderr == ''

    return isSuccess, result.stdout, result.stderr


def splitRunner(
    *,
    splitedCommand: List[str]
):

    result = runnerWrapper(splitedCommand=splitedCommand)
    isSuccess = result.stderr == ''

    return isSuccess, result.stdout, result.stderr
