if ($Args.Length -ne 4) {
    echo "Invalid argument."
    exit 200
}

$remotehost = $Args[0]
$username = $Args[1]
$passwd = $Args[2]
$execution= $Args[3]

$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)
$session = New-PSSession -ComputerName $remotehost -Credential $cred

$stdout = invoke-command -Session $session -ScriptBlock{Invoke-Expression $args[0]} -argumentList $execution
$result = $?

if($result -eq $false){
    exit 250
}

$exitcode = Invoke-Command -Session $session -ScriptBlock { $lastExitCode }

Remove-PSSession -Session $session

echo $stdout

if ($exitcode -eq $null){
    exit 0
}else{
    exit $exitcode
}

exit 0