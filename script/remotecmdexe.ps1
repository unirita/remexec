$remotehost = $Args[0]
$username = $Args[1]
$passwd = $Args[2]
$execution= $Args[3]

if ($Args.Length -ne 4) {
    echo "Invalid argument."
    exit 1
}

$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)
$session = New-PSSession -ComputerName $remotehost -Credential $cred


$stdout = invoke-command -Session $session -ScriptBlock{Invoke-Expression $args[0]} -argumentList $execution
$result = $?

$exitcode = Invoke-Command -Session $session -ScriptBlock { $lastExitCode }

Remove-PSSession -Session $session

echo $stdout

if ($exitcode -eq $null){
    if($result -eq $true){
        echo "ExitStatus: 0"
    }else{
        echo "ExitStatus: 1"
    }
}else{
    echo "ExitStatus: $exitcode"
}

exit 0