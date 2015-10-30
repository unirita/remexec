$commandFlag = $Args[0]
$remotehost = $Args[1]
$username = $Args[2]
$passwd = $Args[3]
$execution= $Args[4]

if ($Args.Length -ne 5) {
    echo "Invalid argument."
    exit 1
}

$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)

$result
if ($commandFlag -eq "e"){
    $result = invoke-command -ComputerName $remotehost -Credential $cred -ScriptBlock{Invoke-Expression $args[0]; echo "ExitStatus: $lastexitcode"} -argumentList $execution
}else{
    $result = invoke-command -ComputerName $remotehost -Credential $cred -FilePath $execution; echo "ExitStatus: $lastexitcode"
}

echo $result

exit 0



